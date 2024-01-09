package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/samber/lo"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

const (
	generatePackageName = "i18n"
	generateDir         = "./generated/" + generatePackageName + "/"
	// 语言定义文件
	languagesDir = "./internal/languages/"
)

// 定义模版文件
//
//go:embed tpl/defs.tmpl
var templateDefines string

//go:embed tpl/languages.tmpl
var templateLanguages string

type LangDefine struct {
	Key  string
	Msg  string
	Args []string
}

type KeyDefine map[string]LangDefine

func main() {
	keys := loadLanguages(languagesDir)
	exportLanguageConstant(keys)
	exportLanguageToMap(keys)
}

var snakeRegex = regexp.MustCompile("_[a-zA-Z]")

func camel(s string) string {
	return snakeRegex.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToUpper(s[1:])
	})
}
func exportable(s string) string {
	return strings.ToUpper(s[:1]) + camel(s[1:])
}

var interpRegx = regexp.MustCompile(`\${\s*(\w+)\s*}`)

// 导出国际化适配的Key值
func exportLanguageConstant(keys map[string]KeyDefine) {
	out, err := os.OpenFile(fmt.Sprintf("%s%s", generateDir, "defs.go"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	t := template.Must(template.New("defines").
		Funcs(template.FuncMap{
			"exportable": exportable,
			"camel":      camel,
			"lastIndex": func(v []string) int {
				return len(v) - 1
			},
		}).
		Parse(templateDefines))

	data := make(map[string]LangDefine, len(keys))
	for k, v := range keys {
		data[k] = v["zh-Hans"]
	}

	if err = t.Execute(out, data); err != nil {
		panic(err)
	}
}

func exportLanguageToMap(keys map[string]KeyDefine) {
	out, err := os.OpenFile(fmt.Sprintf("%s%s", generateDir, "languages.go"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	checkSum := md5.Sum([]byte(fmt.Sprintf("%v", keys)))

	languages := make(map[string]map[string]string)
	for k, defs := range keys {
		for l, v := range defs {
			if languages[l] == nil {
				languages[l] = make(map[string]string)
			}
			languages[l][k] = strings.ReplaceAll(v.Msg, "\n", "\\n")
		}
	}

	tpl := template.Must(template.New("languages").Parse(templateLanguages))
	data := map[string]any{
		"languages": languages,
		"etag":      fmt.Sprintf("%x", checkSum),
		"buildTime": time.Now().Format(time.RFC3339),
	}
	if err := tpl.Execute(out, data); err != nil {
		panic(err)
	}
}

func loadLanguages(dir string) map[string]KeyDefine {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	langDefines := make(map[string]KeyDefine)

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".yaml") {
			continue
		}

		f, _ := os.Open(path.Join(dir, name))
		info, _ := f.Stat()
		if _, err = os.Stat(generateDir); os.IsNotExist(err) {
			if err := os.MkdirAll(generateDir, 0755); err != nil {
				panic(err)
			}
		}

		if !strings.HasSuffix(info.Name(), "yaml") {
			panic("not supported language file type: " + info.Name())
		}

		rawDefs := make(map[string]map[string]string)
		be, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		_ = f.Close()

		if err = yaml.Unmarshal(be, &rawDefs); err != nil {
			panic(err)
		}

		for key, defs := range rawDefs {
			if defs["en"] == "" {
				panic(fmt.Sprintf("key %s, no 'en' language", key))
			}
			if defs["zh-Hans"] == "" {
				panic(fmt.Sprintf("key %s, no 'zh-Hans' language", key))
			}

			if langDefines[key] == nil {
				langDefines[key] = make(KeyDefine, 0)
			} else {
				panic(fmt.Sprintf("key %q duplicated", key))
			}

			for l, v := range defs {
				_, err := language.Parse(l)
				if err != nil {
					panic(fmt.Sprintf("key %s has invalid language define %q", key, l))
				}

				args := lo.Map(interpRegx.FindAllStringSubmatch(v, -1), func(v []string, _ int) string {
					return v[1]
				})

				msg := strings.ReplaceAll(v, "\n\r", " ")
				msg = strings.ReplaceAll(msg, "\n", " ")
				msg = strings.ReplaceAll(msg, "\r", " ")

				if utf8.RuneCountInString(msg) > 30 {
					msg = string([]rune(msg)[:30])
				}

				langDefines[key][l] = LangDefine{
					Key:  key,
					Msg:  v,
					Args: args,
				}
			}
		}
	}

	return langDefines
}

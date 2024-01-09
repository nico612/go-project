// Code generated by global-language. DO NOT EDIT.
package i18n

import (
    "context"
    "fmt"
    "strings"

    "github.com/spf13/cast"
)

// InvalidArgument 参数错误
func InvalidArgument() message {
    return message{ key: "InvalidArgument", args: nil }
}

// PermissionDenied 没有操作权限
func PermissionDenied() message {
    return message{ key: "PermissionDenied", args: nil }
}

// ServerErr 服务器内部错误，错误：${err}
func ServerErr(err any) message {
    args := map[string]string{ "err": cast.ToString(err),  }
    return message{ key: "ServerErr", args: args }
}

// VCodeBlocked 验证码获取太频繁
func VCodeBlocked() message {
    return message{ key: "VCodeBlocked", args: nil }
}

// VCodeEmail 【Telecode】您的验证码：${code}，5分钟内有效，如非本人操作，请忽略本邮件。
func VCodeEmail(code any) message {
    args := map[string]string{ "code": cast.ToString(code),  }
    return message{ key: "VCodeEmail", args: args }
}

// VCodeEmailTitle 邮件验证码
func VCodeEmailTitle() message {
    return message{ key: "VCodeEmailTitle", args: nil }
}

// VCodeInvalid 验证码错误
func VCodeInvalid() message {
    return message{ key: "VCodeInvalid", args: nil }
}

// VCodeSMS 【Telecode】您的验证码：${code}，5分钟内有效，如非本人操作，请忽略本短信。
func VCodeSMS(code any) message {
    args := map[string]string{ "code": cast.ToString(code),  }
    return message{ key: "VCodeSMS", args: args }
}


type message struct {
    key  string
    args map[string]string
}

func (m message) Args() map[string]string {
    return m.args
}

func (m message) Key() string {
    return m.key
}

type Message interface {
    Key() string
    Args() map[string]string
}

type langKey struct{}

func WithLang(ctx context.Context, lang string) context.Context {
    return context.WithValue(ctx, langKey{}, lang)
}

func FromContext(ctx context.Context) string {
    lang := ctx.Value(langKey{})
    if lang == nil {
        return ""
    }

    return lang.(string)
}

func Translate(lang string, msg Message) string {
    m := Languages[lang]
    if m == nil {
        m = Languages["en"]
    }

    content := m[msg.Key()]
    for k, v := range msg.Args() {
        content = strings.ReplaceAll(content, fmt.Sprintf("${%s}", k), cast.ToString(v))
    }

    return content
}

func TranslateFromContext(ctx context.Context, msg Message) string {
    lang := FromContext(ctx)
    return Translate(lang, msg)
}

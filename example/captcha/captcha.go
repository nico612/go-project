package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
)

func main() {
}

// get captcha
func GetCaptch(id string) (string, string) {
	var captchaId string
	var content bytes.Buffer

	if id != "" { // 重新生成验证码
		captcha.Reload(id)
		captchaId = id
	} else {
		// 生成新的验证码
		captchaId = captcha.New() // 默认长度为6
		//captchaId = captcha.NewLen(4) // 指定长度为4
	}

	// 生成图片
	if err := captcha.WriteImage(&content, captchaId, 120, 50); err != nil {
		fmt.Printf("write image error: %v", err)
		return "", ""
	}

	return captchaId, base64.StdEncoding.EncodeToString(content.Bytes())
}

// verify captcha
func VerifyCaptcha(id string, digits string) bool {
	return captcha.VerifyString(id, digits)
}

package main

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

/**

nyaruka/phonenumbers 是 Go 语言中用于处理电话号码的库。它提供了功能强大的工具，用于解析、格式化和验证电话号码，同时还支持国际电话号码的处理。

这个库的主要用途包括：

解析：能够将电话号码字符串解析成可操作的对象，提取出国家代码、区号、本地号码等信息。
格式化：能够将电话号码格式化成特定的格式，符合国际或特定国家/地区的电话号码格式。
验证：可以验证电话号码是否符合特定国家/地区的电话号码规则。

*/

func main() {
	number := "+8618389877234"

	// 解析电话号码
	parsedNumber, err := phonenumbers.Parse(number, "CN")
	if err != nil {
		fmt.Println("解析错误：", err)
		return
	}

	// 获取国家代码
	regionCode := phonenumbers.GetRegionCodeForNumber(parsedNumber)
	fmt.Println("国家代码：", regionCode)
	// 国家代码： CN

	countryCode := parsedNumber.GetCountryCode()
	fmt.Println("country code：", countryCode)
	// 86

	// 格式化电话号码
	formattedNumber := phonenumbers.Format(parsedNumber, phonenumbers.NATIONAL)
	fmt.Println("格式化后的号码:", formattedNumber)
	// 格式化后的号码: 183 8987 7234

	// 验证电话号码
	isValid := phonenumbers.IsValidNumber(parsedNumber)
	fmt.Println("是否有效号码:", isValid)
	// 是否有效号码: true
}

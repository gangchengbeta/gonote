package std

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 字符串常见操作
// Path: std\strings_util.go

// PackageStrings strings包中的函数
func PackageStrings() {
	fmt.Println(strings.Contains("hello", "ll"))
	fmt.Println(strings.Index("hello", "lo"))
	fmt.Println(strings.Replace("hello", "l", "DD", -1))
	fmt.Println(strings.Repeat("hello", 10))
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Fields("helloworld\niam\tgangcheng"))
}

// 中文字符常见操作
func PackageUTF8() {
	str := "狗狼编程"
	fmt.Println(utf8.RuneCountInString("Golang编程"))
	// 判断是否是utf8编码
	fmt.Println(str[:len(str)-3])
	fmt.Println(utf8.ValidString(str[:len(str)-3]))
}

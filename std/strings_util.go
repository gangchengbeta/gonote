package std

import (
	"fmt"
	"strings"
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

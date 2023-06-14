package util

import "fmt"

// Counter 计数器函数
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 每个包都可以有自己的init函数 且可以有多个
// init函数在main函数之前执行
// 执行顺序取决于包的依赖关系
// 被依赖包的全局变量>> 被依赖包的init函数 >> ... >> main包的全局变量>> main包的init函数 >> main函数
var i = 0
var F = func(s string) int {
	fmt.Printf("本次被%s调用\n", s)
	i++
	fmt.Printf("deferUtil被第%v次调用", i)
	return i
}

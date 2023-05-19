package api

import (
	"fmt"
	"strconv"
)

// 这是单行注释
func SayHello(word string) {
	fmt.Printf("\"my go demo\"" + word)
	fmt.Printf("\a")
}

// 变量与常量

// 全局变量
var gv int = 1

// 全局常量
const gc int = 2

// 跨包全局常量
const (
	Version string = "1.0.0"
)

func VarAndConstant() {
	// 变量
	var v1 int //默认值为0
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7)
	// 常量
	const c1 = 1
	const (
		c2 = 8
		c3 = iota // 当前行数 注意行数是从零开始计算的
		c4
		// 默认值为上一行的值
		c5 = 12
		c6
	)
	fmt.Println(c1, c2, c3, c4, c5, c6)
	fmt.Println(gv)
}

// 基本数据类型
func BasicDataType() {
	// 整型
	var (
		n1        = 0b010101
		n2 int8   = 0o72
		n3 uint16 = 0xAF1
	)
	fmt.Printf("n1 = %v, type = %T\n", n1, n1)
	fmt.Printf("n2 = %v, type = %T\n", n2, n2)
	fmt.Printf("n3 = %v, type = %T\n", n3, n3)

	//浮点型
	var (
		f1         = 1.02
		f2 float32 = 1.00002
	)
	fmt.Printf("f1 = %v, type = %T\n", f1, f1)
	fmt.Printf("f2 = %v, type = %T\n", f2, f2)

	// 数值型数据类型转换
	n2 = int8(n3) //(需要注意精度)
	fmt.Printf("n2 = %v, type = %T\n", n2, n2)

	// 字符型
	var (
		c1 byte
		c2      = '0'
		c3 rune = '中'
	)
	fmt.Printf("c1 的码值为 %v ,这个码值对用的字符是 %c  type = %T\n", c1, c1, c1)
	fmt.Printf("c2 的码值为 %v ,这个码值对用的字符是 %c  type = %T\n", c2, c2, c2)
	fmt.Printf("c3 的码值为 %v ,这个码值对用的字符是 %c  type = %T\n", c3, c3, c3)

	// 	布尔型
	var boolean bool = true
	fmt.Printf("boolean = %v, type = %T\n", boolean, boolean)
	// 字符串
	var str string = "hello world"
	fmt.Printf("str = %v, type = %T\n", str, str)
	fmt.Printf(strconv.Itoa(len(str)))
}

// 指针
// * 取值
// & 取地址
func incr(n *int) {
	*n++
	fmt.Println(*n)
}

func Pointer() {
	var src = 2
	incr(&src)
	fmt.Println(src)
}

package api

import "fmt"

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

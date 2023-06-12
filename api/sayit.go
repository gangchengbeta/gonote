package api

import (
	"fmt"
	"gonote/util"
	"strconv"
)

var A = util.F("api.A")

func init() {
	util.F("api.init1")
}
func init() {
	util.F("api.init2")
}

// label & goto (goto语句不推荐使用)
func LabelAndGoto() {
	fmt.Println("\n 标签")
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
	fmt.Println("\n goto")
	fmt.Print("1")
	goto four
	fmt.Print("2")
	fmt.Print("3")
four:
	fmt.Print("4")
}

// for循环
func ForRoll() {
	fmt.Println("\n 无限循环")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 10 {
			fmt.Println()
			break
		}
	}
	fmt.Println("\n 条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println("\n 标准for循环")
	for i := 0; i < 11; i++ {
		fmt.Print(i, "\t")
	}

}

// switchCase
func SwitchDemo(age int) {
	switch {
	case age < 15:
		fmt.Println("小学生")
		fallthrough // 穿透 继续匹配下一项
	case age < 18:
		fmt.Println("中学生")
	case age < 22:
		fmt.Println("大学生")
	default:
		fmt.Println("社会人")
	}
}

// 这是单行注释
func SayHello(word string) {
	fmt.Printf("\"my go demo\"" + word)
	fmt.Printf("\a")
}

// 这是多行注释
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
// * 取值 & 取地址
func incr(n *int) {
	*n++
	fmt.Println(*n)
}

func Pointer() {
	var src = 2
	incr(&src)
	fmt.Println(src)
}

// FmtVerbs fmt格式字符
func FmtVerbs() {
	fmt.Println("\n 通用")
	fmt.Println("%%\n", "百分号")

	fmt.Println("\n 整数")
	i := 123
	fmt.Printf("%U\n", i)
	fmt.Printf("%c\n", i)
	fmt.Printf("%q\n", i)
	fmt.Printf("\n 浮点数")
	f := 123.456
	fmt.Printf("%f\n", f)
	// 保留两位小数
	fmt.Printf("%.2f\n", f)
	// 保留两位小数，宽度为10
	fmt.Printf("%10.2f\n", f)
	// 指数为2的科学计数法(会有不可避免精度误差)
	fmt.Printf("%b", f)

	fmt.Printf("\n 布尔")
	fmt.Printf("%t\n", f == 123456)
	fmt.Printf("\n 字符串或byte切片")
	s := "123"
	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")
	//将每个byte按两位小写十六进制输出
	fmt.Printf("%x\n", "hello world")
	fmt.Printf("\n 指针")
	p := &s
	fmt.Printf("%p\n", p)
}

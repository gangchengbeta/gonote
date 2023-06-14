package note

import (
	"fmt"
)

// 函数
// 函数本质也是数据结构
// 函数名也是一个指针 指向其函数的内存地址
func getSum(n1, n2 int) (sum, difference int) {
	sum = n1 + n2
	difference = n1 - n2
	return
}

// 闭包
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用的参数为%v\n", n)
		i++
		fmt.Printf("deferUtil被第%v次调用", i)
		return i
	}
}

// Defer  延迟执行函数
// 延迟执行的函数会被压入一个栈中 return之后按照先进后出的顺序调用
// 延迟执行的函数其参数会立即求值
// defer常用于资源释放、文件关闭、解锁以及记录时间等操作
func Defer() int {
	f := deferUtil()
	defer f(f(3))
	return f(2)
}
func Funcation() {
	//匿名函数
	var getSum = func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}
	// 立即调用的匿名函数
	res3, res4 := func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	res1, res2 := getSum(10, 20)
	fmt.Println(res1, res2)
	fmt.Println(res3, res4)
	fmt.Printf("getSum = %v , Type of getSum = %T\n", getSum, getSum)
}

// recover 错误捕捉 函数能防止异常错误退出
func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

package note

import (
	"fmt"
	"strings"
)

// LearnString 字符串学习
func LearnString() {
	s := "GOLANG"
	//fmt.Println(s)
	//for i, ele := range s {
	//	fmt.Printf("%d %c\n", i, ele)
	//}
	//// s[0] = `f`   cannot assign to s[0] (value of type byte)
	//fmt.Println(len(s)) // 6
	a := "golang你好"
	//// 一般来说 一个汉字占3个字节
	//fmt.Println(len(a))   // 12
	//arr := []rune(a)      // 将string转换为rune切片 可以求得正确的长度
	//fmt.Println(len(arr)) // 8
	//for _, ele := range a {
	//	fmt.Printf("%c\n", ele)
	//}
	// 字符串的拼接
	// 1. 使用+号
	c := s + a
	fmt.Println(c)
	// 2. 使用fmt.Sprintf
	d := fmt.Sprint(s, a)
	fmt.Println(d)
	// 3. 使用strings.Join
	e := strings.Join([]string{s, a}, "")
	fmt.Println(e)

}

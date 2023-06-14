package std

import (
	"fmt"
	"strconv"
)

// 字符串类型转换
func StrCovert() {
	i := 123
	s1 := "gangcheng"
	s2 := fmt.Sprintf("%d@%s", i, s1)
	fmt.Println(s2)
	var (
		i1 int
		s3 string
	)
	// 将string 按照指定格式转换为其他类型
	// n: 转换成功的参数个数
	// err: 转换失败的错误信息
	n, err := fmt.Sscanf(s2, "%d@%s", &i1, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, i1, s2)
	s4 := strconv.FormatInt(123, 4)
	fmt.Println(s4)
	u1, err := strconv.ParseUint(s4, 4, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(u1)
}

package std

import (
	"fmt"
	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}
type Class struct {
	Id       string    `json:"id"`
	Students []Student `json:"students slice"`
}

func SerialJSON() {
	s := Student{"张宝磊", 58, false}
	c := Class{
		Id:       "1(1)班",
		Students: []Student{s, s, s},
	}
	// json 序列化 通过json.Marshal()函数将一个Go语言中的结构体转换为JSON格式的字符串
	bytes1, err1 := sonic.Marshal(c)
	if err1 != nil {
		panic(err1)
	}
	//for _, v := range bytes {
	//	fmt.Printf("%c", v)
	//}
	str1 := string(bytes1)
	fmt.Println(str1)

	// json 反序列化 通过json.Unmarshal()函数将一个JSON格式的字符串转换为Go语言中的结构体
	var c2 Class
	err1 = sonic.Unmarshal(bytes1, &c2)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("%v", c2)
	// 字节跳动sonic性能更佳
	// sonic 序列化 sonic.Marshal()函数将一个Go语言中的结构体转换为JSON格式的字符串
	bytes, err := sonic.Marshal(c)
	if err != nil {
		panic(err)
	}
	//for _, v := range bytes {
	//	fmt.Printf("%c", v)
	//}
	str := string(bytes)
	fmt.Println(str)

	// sonic 反序列化 sonic.Unmarshal()函数将一个JSON格式的字符串转换为Go语言中的结构体
	var c1 Class
	err = sonic.Unmarshal(bytes, &c1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", c1)
}

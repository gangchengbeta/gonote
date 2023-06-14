package note

import "fmt"

// User 结构体
type User struct {
	Name string `json:"name"` // 结构体标签
	Id   uint32 `json:"id"`
}
type Account struct {
	User
	password string
}
type Contact struct {
	*User
	Remark string `json:"remark"`
}

// Structs 结构体
func Structs() {
	var u1 User = User{
		Name: "张三",
		Id:   111,
	}
	fmt.Printf("u1: %v , Type of u1 : %T \n", u1, u1)

	u1.Id = 222
	fmt.Printf("u1: %v , Type of u1 : %T \n", u1, u1)

	// 结构体指针
	var u2 = &User{
		Name: "李四",
		Id:   333,
	}
	fmt.Printf("u2: %v \n", u2)
	u2.Id = 1000 // (*u2).Id = 1000 的简写
	fmt.Printf("u2: %v \n", u2)
	var a1 = Account{
		User: User{
			Name: "王五",
			Id:   444,
		},
		password: "666",
	}
	fmt.Printf("a1: %v \n", a1)

	var c1 = Contact{
		&u1, "备注",
	}
	fmt.Printf("c1: %v \n", c1)
}

// 自定义类型& 类型别名
func TypeDefine() {
	type mesType uint16
	var textMsg mesType = 1000
	fmt.Printf("textMsg: %v , Type of textMsg : %T \n", textMsg, textMsg)
	// 混用需要类型转换
	var num uint16 = uint16(textMsg)
	fmt.Printf("num: %v , Type of num : %T \n", num, num)

	// 类型别名 只是名字不同 本质上还是原类型
	type myu16 = uint16
	var num2 myu16 = 100
	fmt.Printf("num2: %v , Type of num2 : %T \n", num2, num2)

}

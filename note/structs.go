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

// 在Go语言中，方法是指与特定类型相关联的函数。方法可以访问和操作类型的属性，并且可以在类型的实例上调用。
// 方法的语法如下：
//
//	func (variable_name variable_data_type) function_name() [return_type]{
//		/* 函数体*/
//	}
//
// 参数传递是值传递 如需修改原结构体内容可以通过指针传递
func (u User) printName() {
	u.Id = 1000000
	fmt.Println("name:", u.Name)
}
func (u *User) setId() {
	u.Id = 1000
}
func Method() {
	u := User{
		Name: "张三",
		Id:   111,
	}
	u.setId()
	fmt.Printf("u: %v \n", u)
	u.printName()
	fmt.Printf("u: %v \n", u)

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

// 接口
// 接口是一种类型，它规定了变量有哪些方法。
// 任何类型的方法集中只要拥有接口中的全部方法，那么这个类型就实现了这个接口。
// 接口的定义格式如下：
//
//	type 接口类型名 interface{
//		方法名1( 参数列表1 ) 返回值列表1
//		方法名2( 参数列表2 ) 返回值列表2
//		...
//	}
type textMsg struct {
	Type string
	Text string
}

func (tm *textMsg) setText() {

	tm.Text = "Hello World"
}
func (tm *textMsg) setType() {
	tm.Type = "text"
}

type imgMsg struct {
	Type string
	Img  string
}

func (img *imgMsg) setImg() {
	img.Img = "img.png"
}
func (img *imgMsg) setType() {
	img.Type = "img"
}

type Mes interface {
	setType()
}

// 类型断言
// 类型断言是一个使用在接口值上的操作，语法上它看起来像 x.(T) ，其中 x 是一个接口的类型和 T 是一个类型（也称为断言类型）。
// 还原为原始类型 interface.(Type)
// 如果接口没有包含任何值，那么这个类型断言就会失败，但是如果接口包含了值，那么这个操作就会检查这个接口的值是否是指定的类型。
// 可返回两个值，第一个是值，第二个是是否成功
// value, ok := element.(T)
// 空接口可以接受任何类型的值 interface{}
func SendMsg(mes Mes) {
	mes.setType()
	switch mesPtr := mes.(type) {
	case *textMsg:
		mesPtr.setText()
	case *imgMsg:
		mesPtr.setImg()
	}
	fmt.Printf("mes: %v \n", mes)
}

func Interface() {
	tm := textMsg{}
	SendMsg(&tm)
	img := imgMsg{}
	SendMsg(&img)
}

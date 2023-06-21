package db

// 文档地址 : https://gorm.io/zh_CN/docs/update.html
import (
	"fmt"
	"gorm.io/driver/mysql"
	"strconv"
)
import "gorm.io/gorm"

// User 用户表映射
type User struct {
	Id   int64  `json:"id"`
	Name string `gorm:"column:uname" json:"name"`
}

var client = connectMysql()

// TableName 重写表名
func (User) TableName() string {
	return "t_user"
}
func DBCURD() {
	//user := read()
	//if user == nil {
	//	fmt.Println("表中没有数据")
	//	return
	//}
	//fmt.Printf("读取user: %v\n", user)
	//
	//insert()
	//benchInsert()
	//readAll()
	//update()
	delete()
}
func delete() {
	result := client.Where("length(id) > ?", 2).Delete(&User{}) // 删除id长度大于2的记录
	fmt.Printf("删除的记录数是:%v", result.RowsAffected)
	err := result.Error
	handleError(err)
}

func update() {
	user := User{
		Id: 52222,
		//Name: "张宝磊",
	}
	//err := client.Save(&user).Error
	//handleError(err)
	//err := client.Model(&user).Update("uname", "张宝磊儿子").Error
	//handleError(err)
	err := client.Model(&user).Where("id = ?", 52222).Update("uname", "张堡垒孙子").Error
	handleError(err)
}

// 读取全部数据
func readAll() {
	var users []User
	result := client.Find(&users)
	fmt.Printf("总记录是:%v", result.RowsAffected)
	err := result.Error
	handleError(err)
	fmt.Printf("读取user: %v\n", users)
}

// 读取数据
func read() *User {
	var user User
	err := client.Take(&user).Error
	handleError(err)
	return &user
}

// 插入数据
func insert() {
	user := User{
		Id:   52222,
		Name: "coder4j",
	}
	err := client.Create(&user).Error
	handleError(err)
}

func benchInsert() {
	var users []User
	for i := 0; i < 1000; i++ {
		var user = User{
			Id:   int64(52223 + i),
			Name: "coder" + strconv.Itoa(i),
		}
		users = append(users, user)
	}
	err := client.CreateInBatches(&users, len(users)).Error
	handleError(err)
}

// 连接数据库
func connectMysql() *gorm.DB {
	dataSourceName := "root:coder4jdocker@tcp(112.126.71.240:3340)/db_user?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := gorm.Open(
		mysql.Open(dataSourceName),
		nil,
	)
	handleError(err)
	return client
}

// 处理error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

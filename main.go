package main

import (
	"gonote/note"
)

// 主函数
func main() {
	//engine := gin.Default()
	//engine.GET("/get_name", mygin.GetName)
	//err := engine.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
	//db.RedisCRUD()

	note.CRUD()
}

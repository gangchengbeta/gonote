package main

import (
	"github.com/gin-gonic/gin"
	myclient "gonote/gin"
)

// 主函数
func main() {
	engine := gin.Default()
	engine.GET("/get_name", myclient.GetName)
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

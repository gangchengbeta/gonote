package main

import (
	"gonote/api"
	"gonote/util"
)

var A = util.F("main.A")

func init() {
	util.F("main.init1")
}
func init() {
	util.F("main.init2")
}

// 主函数
func main() {
	api.Funcation()
	//fmt.Println("DeferRecover() 之后仍在执行")
}

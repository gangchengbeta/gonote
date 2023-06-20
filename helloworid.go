package main

import (
	"gonote/note"
)

//var A = util.F("main.A")
//
//func init() {
//	util.F("main.init1")
//}
//func init() {
//	util.F("main.init2")
//}

// 主函数
func main() {
	//api.Map()
	//fmt.Println("DeferRecover() 之后仍在执行")
	//var counter = util.Counter()
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())

	note.Defer()
}

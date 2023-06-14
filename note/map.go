package note

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// Map 键值对
func Map() {

	var m0 map[string]string
	fmt.Println(m0)
	m0 = make(map[string]string, 16)
	fmt.Println(m0)
	m1 := map[string]string{
		"name": "张三",
		"age":  "18",
		"id":   strings.ReplaceAll(uuid.New().String(), "-", ""), // 生成uuid
	}
	fmt.Println(m1)

	// 查找
	v, ok := m1["name1"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	// map的遍历
	for key, value := range m1 {
		fmt.Println(key, "=", value)
	}
	// 删除 通过delete(map,key)函数删除单个元素
	// 删除不存在的元素也不会报错
	delete(m1, "name")
	fmt.Println(m1)
	// 删除所有元素 通过make()函数重新make一次 或者指向nil
	m1 = nil
	m2 := make(map[string]string, 16)
	fmt.Println(m1)
	fmt.Println(m2)
}

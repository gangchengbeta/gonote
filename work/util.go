package work

import "fmt"

func SelectByKey(text ...string) (key int) {
	for i, v := range text {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Println("请输入要选择的序号:")
	scan, err := fmt.Scanln(&key)
	if err != nil {
		return 0
	}
	return scan
}

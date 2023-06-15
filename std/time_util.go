package std

import (
	"fmt"
	"time"
)

// 时间常见操作
func PackageTime() {
	//休眠
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Second * 2)
	}
	fmt.Println()

	// 时间转换
	duration, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration)

}

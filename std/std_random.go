package std

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数
func Random() {
	// 生成0-10的随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%v\n", rand.Intn(10)+1)
}

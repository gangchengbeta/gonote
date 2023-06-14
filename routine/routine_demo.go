package routine

import (
	"fmt"
	"sync"
)

// 判断是否为素数
var (
	c    int = 0
	lock sync.Mutex
)

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%v\t", n)

	lock.Lock()
	c++
	lock.Unlock()

}

func GoRoutine() {
	for i := 2; i < 100001; i++ {
		go PrimeNum(i)
	}
	// 阻塞

	var key string
	_, _ = fmt.Scanln(&key)
	fmt.Printf("\n一共有%v个素数\n", c)
}

// 管道
func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func pushPrimeNum(n int, c chan int) {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	c <- n

}
func Channel() {
	var c1 = make(chan int)
	go pushNum(c1)
	//for {
	//	if v, ok := <-c1; ok {
	//		fmt.Println(v)
	//	} else {
	//		break
	//	}
	//}
	for v := range c1 {
		fmt.Println(v)
	}
	// 带缓冲区
	var c2 = make(chan int, 100)
	for i := 2; i < 100001; i++ {
		go pushPrimeNum(i, c2)
	}
	// 如果不知道channel何时关闭 则可使用select case default 方法
Print: // label标签
	for {
		select {
		case v := <-c2:
			fmt.Println(v)
		default:
			fmt.Println("所有素数已经找到")
			break Print
		}
	}
}

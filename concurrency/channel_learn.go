package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	// 创建管道
	ch := make(chan int, 1)
	fmt.Println(time.Now().Unix())
	// 从管道中读取数据
	go func() {
		a := <-ch
		fmt.Printf("%d 读出%d\n", time.Now().Unix(), a)
	}()
	time.Sleep(2 * time.Second)
	// 向管道中写入数据
	go func() {
		ch <- 4
		fmt.Println(time.Now().Unix(), "写入channel完成")
	}()

	time.Sleep(1 * time.Second)

}

// 生产者消费者
// 目的: 学会使用channel在协程间协调同步
func ProducerConsumer() {
	ch := make(chan int, 100)
	// 生产者
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2) // 2个生产者
	go func() {
		defer waitGroup.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer waitGroup.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// 消费者
	// 因为blockChan是一个无缓冲的通道,所以会阻塞在这里
	// 设置类型为struct{}的原因是因为struct{}不占用内存空间
	blockChan := make(chan struct{}, 0)
	go func() {
		sum := 0
		for {
			a, ojbk := <-ch
			if !ojbk { // 通道关闭&&通道中没有数据 此时会返回false
				break
			}
			sum += a
		}
		fmt.Printf("sum=%d\n", sum)
		blockChan <- struct{}{}
	}()
	waitGroup.Wait()
	close(ch)
	// 会自动阻塞直到往blockChan中存放数据
	<-blockChan
}

func DeadLock() {
	ch := make(chan int, 0)
	//go func() {
	//	ch <- 1
	//	fmt.Println("over")
	//}()

	//time.Sleep(time.Hour) 通过sleep会使得main协程一直阻塞
	// ch <- 2 //通过channel阻塞main 会被系统检测到死锁 从而报错 fatal error: all goroutines are asleep - deadlock!
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	ch <- 1
	//	fmt.Println("over")
	//}()
	//
	//wg.Wait() // 通过waitGroup阻塞main 会被系统检测到死锁 从而报错 fatal error: all goroutines are asleep - deadlock!

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()

	// 新加一个协程
	go func() {
		time.Sleep(time.Hour)
	}()

	wg.Wait() // 此时 系统不会检测到死锁 因为有一个协程在睡眠,系统无法判定在该协程睡眠过后是否会向ch中写入数据/将等待组减一 所以无法检测到死锁
}

func DeadLockSimpleDemo() {
	ch := make(chan int, 0)
	ch <- 1
	//系统会检测到死锁 fatal error: all goroutines are asleep - deadlock!
	fmt.Println("over")
}

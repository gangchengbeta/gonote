package concurrency

import (
	"fmt"
	"gonote/std"
	"sync"
	"sync/atomic"
)

var (
	n int32
	//lock = sync.Mutex{}
)

func add() {
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//n++
		//lock.Unlock()
		// 或者使用并发安全的原子操作 将n++ 强行改为原子操作
		atomic.AddInt32(&n, 1)
	}
	fmt.Printf("add() done. n=%v\n", n)
}
func Export() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		add()
	}()
	go func() {
		defer waitGroup.Done()
		add()
	}()
	waitGroup.Wait()
	fmt.Printf("Export() done n=%v\n", n)
}

// SimpleWaitG WaitGroup 简单使用
func SimpleWaitG() {
	waitGroup := sync.WaitGroup{}
	// 初始化数
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			result, err := std.RedisClient.Ping().Result()
			if err != nil {
				panic(err)
			}
			println(result)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

package concurrency

import (
	"gonote/std"
	"sync"
)

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

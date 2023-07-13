package concurrency

import (
	"context"
	"sync"
	"time"
)

//一道考题
//要求实现一个 map:
//(1) 面向高并发;
//(2) 只存在插入和查询操作 0(1);
//(3)查询时，若 key 存在，直接返回 val;
//			若 key 不存在，阻塞直到 key val 对被放入后，获取 val 返回;
//			等待指定时长仍未放入，返回超时错误;
//(4) 写出真实代码，不能有死锁或者 panic 风险

type ImproveChannel struct {
	sync.Once //确保只关闭一次
	ch        chan struct{}
}

func NewImproveChannel() *ImproveChannel {
	return &ImproveChannel{
		ch: make(chan struct{}),
	}
}

func (channel *ImproveChannel) Close() {
	channel.Once.Do(func() {
		close(channel.ch)
	})
}

type ConcurrentMap struct {
	m          map[string]string
	improveCh  map[string]*ImproveChannel
	sync.Mutex //互斥锁
}

//func NewConcurrentMap() *ConcurrentMap {
//	return &ConcurrentMap{
//		m:         make(map[string]string),
//		improveCh: make(map[string]*ImproveChannel),
//	}
//}

func (concurrentMap *ConcurrentMap) Put(key string, val string) {
	concurrentMap.Lock()
	defer concurrentMap.Unlock()
	concurrentMap.m[key] = val
	channel, ok := concurrentMap.improveCh[key]
	if !ok {
		return
	}
	channel.Close()
}

func (concurrentMap *ConcurrentMap) Get(key string, maxWaitDuration time.Duration) (string, error) {
	concurrentMap.Lock()
	val, ok := concurrentMap.m[key]
	if ok {
		concurrentMap.Unlock()
		return val, nil
	}
	channel, ok := concurrentMap.improveCh[key]
	// 不存在则创建新的channel
	if !ok {
		channel = NewImproveChannel()
		concurrentMap.improveCh[key] = channel
	}
	ctx, cannel := context.WithTimeout(context.Background(), maxWaitDuration)
	defer cannel()
	concurrentMap.Unlock()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-channel.ch:
	}
	concurrentMap.Lock()
	val = concurrentMap.m[key]
	concurrentMap.Unlock()
	return val, nil
}

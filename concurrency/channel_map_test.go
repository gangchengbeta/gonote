package concurrency

import (
	"fmt"
	"testing"
	"time"
)

var myMap = NewConcurrentMap()

func TestPutAndGet(t *testing.T) {
	go myMap.Put("key", "val")
	go myMap.Put("key1", "val1")
	time.Sleep(1 * time.Second)
	go func() {
		value, err := myMap.Get("key", 300*time.Millisecond)
		if err != nil {
			t.Error(err)
		}
		if len(value) == 0 {
			t.Error("无法获取值")

		}
		fmt.Println(value)
	}()
	go func() {
		value, err := myMap.Get("key2", 300*time.Millisecond)
		if err != nil {
			t.Error(err)
		}
		if len(value) == 0 {
			t.Error("无法获取值")

		}
		fmt.Println(value)
	}()
	time.Sleep(1 * time.Second)

}

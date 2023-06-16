package std

import (
	"fmt"
	"math/rand"
)

// 用go实现几种经典排序算法

// Path: std\sort.go
func Sort() {
	var array = make([]int, 100)
	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(1000)
	}
	fmt.Println("排序前: ", array)
	selectionSort(array)
	fmt.Println("排序后: ", array)

}

// 冒泡排序
func bubbleSort(array []int) {
	lastIndex := len(array) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
			if array[j] > array[j+1] {
				// 位置互换
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// selectionSort 选择排序
func selectionSort(array []int) {
	lastIndex := len(array) - 1
	for i := 0; i < lastIndex; i++ {
		max := lastIndex - i // 最大值的下标
		for j := 0; j < lastIndex-i; j++ {
			if array[j] > array[max] {
				max = j
			}
		}
		if max != lastIndex-i {
			// 位置互换
			array[lastIndex-i], array[max] = array[max], array[lastIndex-i]
		}
	}
}

// insertionSort 插入排序 适合将数据插入到已经排好序的数据中
func insertionSort(array []int) {

}

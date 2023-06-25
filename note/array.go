package note

import (
	"fmt"
)

// ArrayDemo 数组
func ArrayDemo() {
	var a = [...]int{1, 119, 120, 110}
	fmt.Println(a)

	//遍历1 for循环
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	//遍历3 for range
	for i, v := range a {
		fmt.Printf("a[%v]: %v\n", i, v)
	}
	// 二维数组
	var twoDimensionalArray = [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	// 遍历二维数组  使用双层遍历
	//下划线是占位符，表示忽略该值
	for _, v := range twoDimensionalArray {
		for _, v2 := range v {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}

}

// SliceDemo 切片
// 切片是对数组的抽象
func SliceDemo() {
	array := [5]int{1, 2, 3, 4, 5}
	// 切片的声明
	// 对于从头开始的切片 可以省略0
	// 对于到尾的切片 可以省略len
	var slice = array[1:3] // 左闭右开区间  [1,3)
	fmt.Println(slice)
	// 对切片的操作会影响到原数组
	slice[0] = 100
	fmt.Println(array)
	// 可以对切片再次切片
	slice2 := slice[1:2]
	fmt.Println(slice2)
	slice2[0] = 200
	fmt.Println(array)
	var slice4 []int // 创建但是没有分配内存空间的切片
	fmt.Println("does slice4 == nil?", slice4 == nil)
	// 如果想要切片不影响原数组 可以使用copy函数
	// 如果不需要原数组 可以直接使用make函数使系统分配内存空间
	// make([]T, len, cap)
	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)
	// 切片的长度和容量
	fmt.Printf("len = %v, cap = %v\n", len(slice3), cap(slice3))
	// 切片的扩容
	// 当切片的容量不够时，系统会自动扩容
	// 底层原理：切片扩容时会重新分配一个数组，将原数组的值拷贝到新数组中
	// append([]T, ...T)
	slice5 := append(slice3, 1, 2, 3, 4, 5, 6, 7, 8)
	slice5[0] = 100
	fmt.Println(slice5)
	fmt.Println(slice3)
	// 复制数组
	// copy([]T, []T)
	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := make([]int, 5, 5)
	copy(slice7, slice6) // 容量不够时，只会复制最大容量的值
	fmt.Println("slice6:", slice6)
	fmt.Println("slice7:", slice7)
	// string与byte切片
	// string底层是byte数组 可以项目转换
	// 个数字符通用
	// 可以直接用for range遍历字符串
	str := "nihaome imissyou"
	fmt.Printf("[]byte(str): %s\n", []byte(str))

	for i, v := range str {
		fmt.Printf("str[%v]: %c\n", i, v)
	}
}

func AppendDemo() {
	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 2, 7, 9
	//底层数组 : 2, 7, 9, 8
	brr := append(arr, 8)
	// brr和arr共享底层数组 因此修改arr会影响brr 反之亦然
	fmt.Println(len(arr)) // 3
	fmt.Println(len(brr)) // 4
	//底层数组 : 2, 7, 9, 10
	arr = append(arr, 10)
	fmt.Println(arr)
	fmt.Println(brr)
	fmt.Println("brr 插入 	容量满")
	brr = append(brr, 15)
	arr[0] = 100
	brr[1] = 200
	fmt.Println("容量满 arr", arr) //容量满 arr [100 200 9 10]
	fmt.Println("容量满 brr", brr) //容量满 brr [100 200 9 10 15]
	//根据以上运行结果可知 当容量未满时，arr和brr始终共享底层数组

	fmt.Println("brr 插入  需要扩容")
	brr = append(brr, 16)
	arr[2] = 200
	brr[3] = 300
	fmt.Println("需要扩容 arr", arr) //需要扩容 arr [100 200 200 10]
	fmt.Println("需要扩容 brr", brr) //需要扩容 brr [100 200 9 300 15 16]
	// 根据以上运行结果可知 当容量满时，arr和brr不再共享底层数组
	// 底层会重新分配一个数组，将原数组的值拷贝到新数组中 将新数组的地址赋值给brr的Pointer 从而对brr的修改不会影响arr
}

func CRUD() {
	var sliceName []string
	// 增加
	sliceName = append(sliceName, "张三", "李四", "王五")
	fmt.Println(sliceName)
	// 修改
	sliceName[0] = "赵六"
	fmt.Println(sliceName)
	// 删除
	sliceName = append(sliceName[:1], sliceName[2:]...)
	fmt.Println(sliceName)
	// 查找
	for i, v := range sliceName {
		fmt.Printf("sliceName[%v] = %v\n", i, v)
	}
}

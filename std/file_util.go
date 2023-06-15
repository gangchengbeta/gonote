package std

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 文件操作
// 判断文件是否存在
func fileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir()
	}
	return os.IsExist(err)
}

// 根据文件路径创建文件夹
func createDir(path string) error {
	paths := strings.Split(path, "/")
	fmt.Println(paths)
	var dirPath = ""
	for i, v := range paths {
		if i == len(paths)-1 {
			break
		}
		if i != 0 {
			dirPath += "/"
		}
		dirPath += v
	}
	fmt.Println(dirPath)
	return os.MkdirAll(dirPath, os.ModePerm)
}
func FileOperation() {
	//dir := readDir("D:\\CodeProjects")
	//fmt.Println(dir)
	fileRW("2d1.txt")
}

// 读取文件夹下的所有文件
func readDir(path string) []string {
	var file = make([]string, 16)
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		file = append(file, entry.Name())
	}
	return file
}

// 文件读写
func fileRW(path string) {
	// 判断文件是否存在
	if !fileExist(path) {
		// q: go中如何创建文件
		// a: os.Create(path) 会创建文件
		_, err := os.Create(path)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	// 写入文件
	_, err = writer.WriteString("hello world")
	if err != nil {
		panic(err)
	}
	// 刷新缓冲区
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

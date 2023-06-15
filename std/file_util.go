package std

import (
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
	//err := createDir("D:/test/1/text.md")
	//if err != nil {
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	dir := readDir("D:\\CodeProjects")
	fmt.Println(dir)
}

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

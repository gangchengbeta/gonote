package std

import (
	"bufio"
	"fmt"
	"io"
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
	fileRWBuf("d1.txt")
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
	// 打开文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) //os.O_TRUNC 清空文件 os.O_APPEND 追加写入
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

	//读文件
	readByteSlice := make([]byte, 1024)
	read, err := file.Read(readByteSlice)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readByteSlice[:read]))
	//写入文件
	context := "大多大事单倍行距阿萨德大萨达是多少收到是是"
	write, err := file.WriteString(context)
	if err != nil {
		panic(err)
	}
	fmt.Println(write) // 成功向文件中写入了多少字节

	//writer := bufio.NewWriter(file)
	//_, err = writer.WriteString("hello world")
	//if err != nil {
	//	panic(err)
	//}
	// 刷新缓冲区
	//err = writer.Flush()
	//if err != nil {
	//	panic(err)
	//}
}
func fileRWBuf(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// 延时关闭文件 (必须有)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	//使用缓冲区读取文件 减少磁盘IO提高程序性能
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 碰到换行符就结束读取

		if err != nil {
			// 判断是否读取到文件末尾
			if err == io.EOF {
				fmt.Println(line)
				fmt.Println("文件读取完毕")
				break
			}
			panic(err)
		}
		fmt.Print(line)
	}

	//使用缓冲区写入文件 减少磁盘IO提高程序性能
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("hello world")
	if err != nil {
		panic(err)
	}
	//刷新缓冲区
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

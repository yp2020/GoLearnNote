package file

import (
	"fmt"
	"os"
)

func LearnFile() {
	readFile()
	writeFile()
}

func readFile() {
	// 打开文件
	// 为什么这里的 a.txt 要在 main.go 的同一级目录下呢？
	// 因为 生成的 2进制文件 main，是相对于这个二进制文件,而不是相对于当前文件夹的
	file, err := os.Open("./file/a.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 需要保证文件能够成功打开，才关闭
	defer file.Close()
	// 读取出来的文件内容需要一个字节切片来接收
	var tmp = make([]byte, 128)
	// 读取文件内容
	// 读文件
	nums, err := file.Read(tmp)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println("读取了", nums, "个字节数据")
	fmt.Println(tmp[:nums])
	fmt.Println(tmp)
	fmt.Println(string(tmp[:nums]))
	// 关闭文件
	// 确保文件一定能关闭, 用defer
}

func writeFile() {
	// 写文件
	//打开文件用 OpenFile
	// 如果文件不存在，则创建文件 os.O_CREATE
	// 如果文件存在，则清空文件 os.O_TRUNC 追加写 os.O_APPEND
	// 定义文件的权限 os.O_WRONLY
	file, err := os.OpenFile("./file/b.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello world"
	// 写入文件
	n, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	fmt.Println("写入了", n, "个字节数据")
}

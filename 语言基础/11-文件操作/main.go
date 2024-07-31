package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "D:\\code\\goprojects\\11-文件操作\\example.txt"
	//0666权限，可读可写可执行
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close() //确保文件在函数结束时关闭

	//读取内容
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("read file failed, err:", err)
	} else {
		fmt.Println("read file success, content:", string(content[:n]))
	}

}

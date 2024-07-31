package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.OpenFile(".\\test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	defer file.Close()
	// 写入文件
	content := "hello world\n"
	n, err := file.WriteString(content)
	if err != nil {
		fmt.Println("write string failed, err:", err)
	} else {
		fmt.Println("write string success, n:", n)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line) // 输出最后一行内容
				break
			}
		} else {
			fmt.Print(line) // 输出每一行内容
		}
	}

}

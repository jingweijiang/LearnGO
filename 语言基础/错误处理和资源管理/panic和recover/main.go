package main

import "fmt"

func callThrowPanic() {
	defer func() {
		if recover() != nil {
			println("recover from panic")
		}
	}()

	defer println("defer 1")
	throwPanic() // 这里panic，会立即终止函数正常流程，执行注册的defer方法
	fmt.Println("this line will not be executed")
}

func throwPanic() {

	panic("this is a panic")
}

func main() {
	callThrowPanic()
	fmt.Println("main function end")
}

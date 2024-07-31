package main

import (
	"fmt"
	"time"
)

func f1() {

	//defer recover()	// 不会生效

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("f1 panic:", err)
		}
	}()
	a, b := 1, 0
	_ = a / b
	fmt.Println("f1 finished")
}

func main() {
	go f1()
	time.Sleep(3 * time.Second)
	fmt.Println("main")
}

package main

import (
	"fmt"
	"time"
)

var done chan bool = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello world"
	//done <- true // 通知 main goroutine 完成
	close(done) // 关闭 done 通道
}

var limit = make(chan int, 3)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second) },
	func() { println("2"); time.Sleep(1 * time.Second) },
	func() { println("3"); time.Sleep(1 * time.Second) },
	func() { println("4"); time.Sleep(1 * time.Second) },
	func() { println("5"); time.Sleep(1 * time.Second) },
}

func main() {
	// 启动一个 goroutine
	//go aGoroutine()
	//<-done // 等待 goroutine 完成
	//println(msg)

	for _, w := range work {
		fmt.Println("start", w)
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select {}
}

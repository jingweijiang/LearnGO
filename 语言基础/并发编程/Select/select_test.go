package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		time.Sleep(time.Second * 2)
	//		fmt.Println("hello")
	//		ch <- 1
	//	}()
	//}
	// 基于 select 实现的管道的超时判断
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
		return
	}

}

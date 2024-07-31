package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var lock sync.Mutex

func foo() {
	for i := 0; i < 1000000; i++ {
		//通过加锁来实现数据安全
		//lock.Lock()	//方式一
		//counter++
		//lock.Unlock()
		atomic.AddInt32(&counter, 1) //方式二
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		foo()
	}()
	go func() {
		defer wg.Done()
		foo()
	}()
	wg.Wait()
	fmt.Println("main: ", counter)
}

package main

import (
	"fmt"
	"sync"
)

// 2个生产者，1个消费者

func main() {

	var sum int

	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int, 0)
	// 生产者1
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	//wg1 := sync.WaitGroup{}
	//wg1.Add(1)

	ch2 := make(chan struct{}, 0) // 协程之间协调作用，用于通知生产者结束
	// 消费者
	go func() {
		//defer wg1.Done()
		for {
			i, ok := <-ch
			if !ok {
				ch2 <- struct{}{}
				break
			} else {
				sum += i
			}

		}
	}()

	// 等待生产者处理结束
	wg.Wait()
	// 关闭channel不再接收数据
	close(ch)
	// 等待消费者处理结束
	//wg1.Wait()
	<-ch2 // 等待消费者处理结束
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("foo")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("bar")
	}()
	wg.Wait()
	fmt.Println("main")

}

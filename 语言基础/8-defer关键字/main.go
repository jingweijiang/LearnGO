package main

import "fmt"

func foo() int {

	a, b := 1, 2
	c := a + b
	defer fmt.Println("111", c) // 111 3,

	fmt.Println("222", c) // 222 3

	defer fmt.Println("333", c) // 333 3

	defer func() {
		fmt.Println("444", c) // c = 100,匿名函数执行时，c的值为100
	}()

	c = 100
	return c

}

func main() {
	foo()

}

package main

import "fmt"

func f(a int) {
	fmt.Println(a)
}

func main() {
	a := 10
	go f(a)
	fmt.Println("main...")
}

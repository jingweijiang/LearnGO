package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Even")
			continue
		} else if i > 20 {
			break
		}
	}
}

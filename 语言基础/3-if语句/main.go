package main

import "fmt"

func main() {

	a, b := 10, true

	if a > 0 {
		fmt.Println("a > 10")
	} else if b {
		fmt.Println("b is true")

	} else {
		fmt.Println("final")
	}

}

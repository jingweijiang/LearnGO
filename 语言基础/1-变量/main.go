package main

import "fmt"

func main() {

	var a int
	var b int
	var c int

	var age byte
	var sex bool
	var price float64

	a = 10
	b = 20
	c = a + b

	f := 40.0
	g := c - a

	fmt.Println(a, b, c)

	fmt.Printf("%d %t %f\n", age, sex, price)
	fmt.Printf("%T %T %T\n", age, f, g)

}

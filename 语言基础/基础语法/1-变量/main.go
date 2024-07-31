package main

import "fmt"

func foo() (int, string) {
	return 1, "a"
}

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

	// 标准声明
	var name string = "alex"
	var age1 int = 2
	fmt.Println(name, age1)

	// 一次声明多个
	var q, w = 1, 2
	fmt.Println(q, w)

	// 声明多个
	var (
		e string
		r int
		t bool
		y float32
	)

	fmt.Println(e, r, t, y)

	// 匿名变量
	x, _ := foo()
	fmt.Println(x)

}

package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i // 0
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Println(b)
	fmt.Println(*b) // 指针取值（根据指针去内存取值）

	//new
	c := new(int) // 申请一块内存，返回指针
	d := new(bool)
	fmt.Println(*c) // 指针取值（根据指针去内存取值）
	fmt.Println(*d)

	// make用于slice、map以及channel的内存创建
	makeSlice := make([]int, 5)
	fmt.Println(makeSlice)

	p := 1
	incr(&p)
	fmt.Println(p)

	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

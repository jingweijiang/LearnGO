package main

import "fmt"

// 单个声明
const name = "3.14"

// 批量声明
const (
	a = 1
	b = 2
)

// 批量声明（简略版）
const (
	n1 = 1000
	n2
	n3
)

const (
	x = iota
	y
	z
)

const (
	p = iota
	o = 100
	i = iota
	u
)

func main() {
	fmt.Println(p)
	fmt.Println(o)
	fmt.Println(i)
	fmt.Println(u)
}

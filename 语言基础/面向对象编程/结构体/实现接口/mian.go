package main

import "fmt"

type baseHuman struct {
	name  string
	age   int
	phone string
}

type Student struct {
	baseHuman // 继承了 baseHuman 结构体 匿名字段
	school    string
	loan      float32
}

type Employee struct {
	baseHuman // 继承了 baseHuman 结构体 匿名字段
	company   string
	salary    float32
}

func (h *baseHuman) SayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old, my phone number is %s.\n", h.name, h.age, h.phone)
}

func (h *baseHuman) Sing(lyric string) {
	fmt.Println("la la la la la la", lyric)
}

func (e *Employee) SayHello() {
	fmt.Printf("Hi, I am %s, i work at %s, call me %s \n", e.name, e.company, e.phone)
}

type Human interface {
	SayHello()
	Sing(lyric string)
}

func main() {
	mike := Student{baseHuman{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	tom := Employee{baseHuman{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定义接口Human类型的变量i
	var i Human
	i = &mike
	i.SayHello()
	i.Sing("I love you")

	i = &tom
	i.SayHello()
	i.Sing("I love you too")
}

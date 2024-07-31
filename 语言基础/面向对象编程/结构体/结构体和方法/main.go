package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name   string
	Age    int
	Height float32
	Sex    bool
}

// 给结构图添加方法
func (h Human) SayHello() {
	fmt.Println("Hello, my name is", h.Name)
}

// 使用值类型，那么方法内对这个值的修改不会生效
func (h Human) SetAge(age int) {
	h.Age = age
}

// 通过type 声明其他类型

type IntSlice []int

func (s IntSlice) Index(i int) int {
	return s[i]
}

// 字段标签
// encoding/json 包使用json标签来控制结构体的编解码行为
// orm 包使用db标签来映射数据库字段
// validate 包使用validate标签来验证字段值
type Node struct {
	Value int   `json:"value" db:"value_column" validate:"required,min=0"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func main() {

	n := Node{
		Value: 5,
		Left: &Node{
			Value: 10,
		},
		Right: &Node{
			Value: 3,
		},
	}
	byts, _ := json.Marshal(n)
	fmt.Println(string(byts))

	a := Human{
		Name:   "alex",
		Age:    20,
		Height: 1.73,
		Sex:    false,
	}

	fmt.Println(a)
	fmt.Printf("%d %s %.2f %t \n", a.Age, a.Name, a.Height, a.Sex)
	fmt.Printf("%v \n", a)
	fmt.Printf("%+v \n", a)
	fmt.Printf("%#v \n", a)
	a.SayHello()

	a.SetAge(25)
	fmt.Println(a)

}

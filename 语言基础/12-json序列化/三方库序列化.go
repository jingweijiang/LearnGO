package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

type NewStudent struct {
	Name string
	Age  int

	Gender bool
}

type NewClass struct {
	ID         string
	NewStudent []NewStudent
}

func main() {
	// TODO: 实现json序列化
	// 1. 创建一个Student实例
	student2 := NewStudent{
		Name:   "Alice",
		Age:    20,
		Gender: true,
	}
	// 1. 创建一个Class实例
	class1 := NewClass{
		ID:         "1001",
		NewStudent: []NewStudent{student2, student2},
	}
	// 2. 将Class实例序列化为json字符串
	jsonStr, err := sonic.Marshal(class1)
	if err != nil {
		fmt.Println("json序列化失败", err)
	} else {
		fmt.Println("json序列化成功", string(jsonStr))
	}

	// 3. 将json字符串反序列化为Class实例
	var class2 NewClass
	err = sonic.Unmarshal(jsonStr, &class2)
	if err != nil {
		fmt.Println("json反序列化失败", err)
	} else {
		fmt.Println("json反序列化成功", class2)
	}
}

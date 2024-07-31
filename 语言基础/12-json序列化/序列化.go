package main

//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Student struct {
//	Name string
//	Age  int
//
//	Gender bool
//}
//
//type Class struct {
//	ID       string
//	Students []Student
//}
//
//func main() {
//	// TODO: 实现json序列化
//	// 1. 创建一个Student实例
//	student1 := Student{
//		Name:   "Alice",
//		Age:    20,
//		Gender: true,
//	}
//	// 1. 创建一个Class实例
//	class1 := Class{
//		ID:       "1001",
//		Students: []Student{student1, student1, student1},
//	}
//	// 2. 将Class实例序列化为json字符串
//	jsonStr, err := json.Marshal(class1)
//	if err != nil {
//		fmt.Println("json序列化失败", err)
//	} else {
//		fmt.Println("json序列化成功", string(jsonStr))
//	}
//
//	// 3. 将json字符串反序列化为Class实例
//	var class2 Class
//	err = json.Unmarshal(jsonStr, &class2)
//	if err != nil {
//		fmt.Println("json反序列化失败", err)
//	} else {
//		fmt.Println("json反序列化成功", class2)
//	}
//}

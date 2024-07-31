package main

import "fmt"

func main() {
	s := "golang你好"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//golang长度是6
	// 汉字每个字符占用3个字节
	fmt.Printf("%d\n", len(s))

	sr := []rune(s) // 将字符串转换为[]rune类型
	fmt.Println(sr)
	for _, v := range sr {
		fmt.Printf("%c\n", v)
	}

	// 字符串拼接
	s1 := "hello"
	fmt.Println(s1)
	s2 := "world"
	fmt.Println(s2)
	s3 := "hello" + " " + "world"
	fmt.Println(s3)

	// 单引号字符串
	s4 := `
	hello
	world
	`
	fmt.Println(s4)
}

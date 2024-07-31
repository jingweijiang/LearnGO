package main

import "fmt"

// 导入包
// 包名：
// import "github.com/gopherchina/gopherchina.org
// 包起别名：
// import gopher "github.com/gopherchina/gopherchina.org"

// 导入包的init函数
// 包的init函数在包被导入时，只会被执行一次。
// import _ "github.com/gopherchina/gopherchina.orgt"

// 可见性
// 包内的函数、变量、常量默认是可见的，但如果要限制对外的访问权限，可以用大写字母开头的名称来隐藏。
// 例如：
// var myPrivateVariable int // 私有变量
// func myPrivateFunction() {} // 私有函数
// const MY_PRIVATE_CONST = 10 // 私有常量

func init() {
	fmt.Println("init 1")
}
func init() {
	fmt.Println("init 2")
}

type Node struct {
}

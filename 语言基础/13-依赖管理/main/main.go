package main

import (
	"fmt"
	"g6/util"
	"g6/util/math"
	//math "g6/util/math"	// 也可以这样导入重命名包
	//_ "g6/util/math"	// 也可以这样导入包，但不使用包内的任何变量，只执行导入包的初始化init函数
	"github.com/bytedance/sonic"
)

func main() {
	fmt.Println(util.Name)
	fmt.Println(util.Add(1, 2))
	fmt.Println(maths.Add(1, 2, 3))
	//fmt.Println(maths.sub(1, 2))
	bytes, _ := sonic.Marshal("hello, world")
	fmt.Println(string(bytes))
}

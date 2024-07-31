package main

import "fmt"

func main() {
	//case1 初始化数组时可以使用初始化列表来设置数组元素的值
	var arr [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

	var arr3 [3]string = [3]string{"hello", "world", "go"}

	//case2 我们可以让编译器根据初始值的个数自行推断数组的长度
	var arr4 [5]int = [...]int{1, 2, 3, 4, 5}

	//case3 使用指定索引值的方式来初始化数组
	var arr5 [5]int = [...]int{1: 10, 3: 30, 4: 40}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
}

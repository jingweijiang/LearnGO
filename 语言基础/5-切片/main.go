package main

import "fmt"

func main() {

	//长度是3，容量是5
	arr := make([]int, 3, 5)

	arr[0], arr[1], arr[2] = 2, 7, 9

	brr := arr
	brr[0] = 22 // 底层公用
	brr = append(brr, 88)
	arr = append(arr, 55)
	fmt.Println(arr)
	fmt.Println(brr)

	for i, e := range arr {
		fmt.Println(i, e)
	}

}

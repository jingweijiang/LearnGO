package main

import "fmt"

// 切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片

func case1() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Printf("s: %v, len(s): %d, cap(s): %d\n", s, len(s), cap(s))
}

func case2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

func main() {
	case2()
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

package main

import "fmt"

func f1() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func f2() {
	//for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func f3() {
	// for循环的初始语句和结束语句都可以省略, 等价于其他开发语言的while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func f4() {

	i := 0
	for {
		i++
		fmt.Println("loop")
		if i > 10 {
			break
		}
	}
}

func f5() {
	//case 1 遍历数组
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func f6() {
	//case 2 遍历map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func f7() {
	//case 3 遍历channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Even")
			continue
		} else if i > 20 {
			break
		}
	}
}

package main

import "fmt"

func case1() {
	//make初始化
	scoreMap := make(map[string]int, 8)
	scoreMap["Alice"] = 90
	scoreMap["Bob"] = 80
	fmt.Println(scoreMap)
}

func case2() {
	//字面量初始化
	userInfo := map[string]string{
		"name":  "jack",
		"email": "jackjack@gmail.com",
	}

	fmt.Println(userInfo)
}

func case3() {
	// 判断key是否存在
	scoreMap := map[string]int{
		"Alice": 90,
		"Bob":   80,
	}

	if _, ok := scoreMap["Alice"]; ok {
		fmt.Println("Alice's score is", scoreMap["Alice"])
	} else {
		fmt.Println("Alice's score is not found")
	}

}

func case4() {
	// for range遍历
	scoreMap := map[string]int{
		"Alice": 90,
		"Bob":   80,
	}

	//for name, score := range scoreMap {
	for name := range scoreMap {
		//fmt.Println(name, ":", score)
		fmt.Println(name)
	}
}

func case5() {
	// 删除元素
	scoreMap := map[string]int{
		"Alice": 90,
		"Bob":   80,
	}

	delete(scoreMap, "Alice")
	fmt.Println(scoreMap)
}

func main() {
	case5()
}

package __1switch

import "fmt"

func switch1() {
	//一个case分支，只有一个值

	finger := 3
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:

		fmt.Println("Index Finger")
	case 3:
		fmt.Println("Middle Finger")
	}
}

func switch2() {
	//一个case分支，有多个值

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:

		fmt.Println("偶数")
	default:
		fmt.Println("Middle Finger")
	}
}

func switch3() {
	//一个case分支，可以用表达式
	age := 20
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age >= 18 && age < 60:
		fmt.Println("成年")
	case age >= 60:
		fmt.Println("老年")
	default:
		fmt.Println("活着真好")
	}
}

func switch4() {
	// fallthrough语法可以执行满足条件的case的下一个case
	s := "apple"
	switch {
	case s == "apple":
		fmt.Println("苹果")
		fallthrough
	case s == "banana":
		fmt.Println("香蕉")
	default:

		fmt.Println("其他水果")
	}
}

func main() {
	switch4()
}

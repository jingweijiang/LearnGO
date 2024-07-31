package main

import (
	"errors"
	"fmt"
)

func add(a, b int) (int, error) {
	return a + b, nil
}
func sub(a, b int) (int, error) {
	return a - b, nil
}

func mul(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数为0")
	}
	return a * b, nil
}

func main() {

	a, b := 1, 0
	_, err := mul(a, b)
	if err == nil {
		fmt.Println("商为：0")
	} else {
		fmt.Println("发生了异常：", err)
	}

}

package util

import "fmt"

var (
	Name = "大脸猫"
)

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("init util")

}

package 小技巧

import (
	"fmt"
	"runtime"
	"testing"
)

// 0 代表当前函数的调用者，
// 1 代表调用者的调用者，
// 2 代表调用者的调用者的调用者，

func af() (string, int) {
	_, filename, line, _ := runtime.Caller(2)
	return filename, line
}

func bf() (string, int) {
	return af()
}

func TestCaller(t *testing.T) {
	filename, line := bf()
	fmt.Println(filename, line)
}

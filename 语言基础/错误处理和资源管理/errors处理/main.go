package main

import (
	"errors"
	"fmt"
)

func getErrors() error {
	// 创建一个简单的error值
	return errors.New("this is an error")
}

type bizError struct {
	code int32
	msg  string
}

func (e *bizError) Error() string {
	return fmt.Sprintf("code : %d, msg : %s", e.code, e.msg)
}

func main() {

	err := getErrors()
	fmt.Println("err:", err)

	// 如果具体的err值实现了Unwrap() error方法，会执行该方法，拿到内部的err
	werr := fmt.Errorf("%w :", err)
	fmt.Println(werr.Error())
	fmt.Println(errors.Unwrap(werr).Error())
	fmt.Println(errors.Is(werr, err))

	var be bizError
	if errors.As(werr, &be) {
		fmt.Sprintf("code : %d, msg : %s", be.code, be.msg)
	}

}

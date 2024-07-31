package main

import (
	"errors"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// case 1 常规写法
func add1(x int, y int) int {
	return x + y
}

// case 2 省略写法（如果多个参数类型一致的话）
func add2(a, b int) (int, error) {
	return a + b, nil
}

// case 3 可变参数
func add3(y ...int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}

// case 1 多返回值
func swap(x, y string) (string, string) {
	return y, x
}

// case 2 返回值命名
func sub(a, b int) (c int, e error) {
	c = a - b
	return c, nil
}

func mul(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数为0")
	}
	return a * b, nil
}

// 函数作为参数
func NewFunc(x, y int) int {
	return x + y
}

func suber(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return suber, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}

}

// defer语句
func deferFunc() {
	fmt.Println("开始执行")
	defer fmt.Println("执行结束 1")
	fmt.Println("继续执行")
	defer fmt.Println("执行结束 2")
}

// 内置函数
// close(ch) 关闭channel
// len(s) 字符串长度
// make(t, args) 创建切片、map、channel
// new(t) 创建指针
// cap(s) 切片容量
// copy(dst, src) 复制切片
// append(s, args) 追加切片元素
// delete(m, key) 删除map元素
// complex(real, imag) 创建复数
// real(c) 复数的实部
// imag(c) 复数的虚部
// panic(v) 引发异常
// recover() 捕获异常并返回错误信息

func main() {

	a, b := 1, 0
	_, err := mul(a, b)
	if err == nil {
		fmt.Println("商为：0")
	} else {
		fmt.Println("发生了异常：", err)
	}

	//匿名函数
	nonFunc := func(x, y int) int {
		return x + y
	}

	fmt.Println(nonFunc(1, 2))

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	adder := adder()
	fmt.Println(adder(40))
	fmt.Println(adder(50))

	deferFunc()
}

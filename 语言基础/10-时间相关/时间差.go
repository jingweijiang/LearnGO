package main

import (
	"fmt"
	"time"
)

const (
	DATE_FORMAT = "2006-01-02"
	TIEM_FORMAT = "2006-01-02 15:04:05"
)

func main() {
	t0 := time.Now()
	fmt.Println(t0.Unix())
	time.Sleep(1 * time.Second)
	t1 := time.Now()
	diff := t1.Sub(t0)
	fmt.Println(diff.Milliseconds()) // 毫秒

	fmt.Println(time.Since(t0).Milliseconds())

	d := time.Duration(2 * time.Second)
	t2 := t0.Add(d)
	fmt.Println(t2.Unix())

	// 时间戳格式化
	t3 := time.Now()
	fmt.Println(t3.Format(DATE_FORMAT))
	fmt.Println(t3.Format(TIEM_FORMAT))

	// 字符串转化为时间
	loc, _ := time.LoadLocation("Asia/Shanghai") // 设置时区
	t4, _ := time.ParseInLocation(TIEM_FORMAT, t3.Format(TIEM_FORMAT), loc)
	fmt.Println(t4.Unix())

}

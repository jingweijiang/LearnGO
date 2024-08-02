package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

type BenchmarkStudent struct {
	Name string
	Age  int

	Gender bool
}

type BenchmarkClass struct {
	ID               string
	BenchmarkStudent []BenchmarkStudent
}

var (
	s1 = BenchmarkStudent{
		Name:   "Tom",
		Age:    20,
		Gender: true,
	}
	c1 = BenchmarkClass{
		ID:               "1001",
		BenchmarkStudent: []BenchmarkStudent{s1, s1, s1, s1}}
)

// go test -bench=Json .\json_benchmark_test.go
//
//	BenchmarkJson-16          350386（一共运行次数）              3461 ns/op(每次运行消耗的时间)
func BenchmarkStdJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonStr, _ := json.Marshal(c1)
		var c3 BenchmarkClass
		_ = json.Unmarshal(jsonStr, &c3)

	}

}

// benchmem 显示内存分配情况
// go test -bench=Sonic .\json_benchmark_test.go -benchmem
// BenchmarkSonic-16        1344217               879.2 ns/op
func BenchmarkSonicJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonStr, _ := sonic.Marshal(c1)
		var c3 BenchmarkClass
		_ = sonic.Unmarshal(jsonStr, &c3)

	}
}

// 基准测试命令
// go test -bench=Json .\json_benchmark_test.go
// go test -bench=Sonic .\json_benchmark_test.go -benchmem
// go test -run=^$ -bench=Json -benchmem -count=1 -benchtime=3s -cpuprofile=cpu -memprofile=mem -blockprofile=block .\json_benchmark_test.go
// -run	运行指定函数
// -bench	运行指定函数的基准测试
// -count	运行次数
// -timeout	超时时间
// -v		详细输出
// -benchtime 3s 运行时间
// -benchmem 内存分配情况
// -cpuprofile cpu性能分析
// -memprofile mem性能分析
// -blockprofile block性能分析

// 性能分析
// go tool pprof  cpu
// go tool pprof  mem
// go tool pprof  block

// top 显示函数调用关系
// list 函数名 显示函数代码
// web 显示函数调用关系的图形化界面

//	goos: windows   操作系统
//	goarch: amd64    处理器架构
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz    处理器型号
// 16核 32线程
// 350386 一次运行的次数
// 3461 ns/op 每次运行消耗的时间
// 906 B/op 内存分配情况
//10 allocs/op
// BenchmarkSonic-16        1250222               964.1 ns/op           906 B/op         10 allocs/op

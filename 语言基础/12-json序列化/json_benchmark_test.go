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
func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonStr, _ := json.Marshal(c1)
		var c3 BenchmarkClass
		_ = json.Unmarshal(jsonStr, &c3)

	}

}

// benchmem 显示内存分配情况
// go test -bench=Sonic .\json_benchmark_test.go -benchmem
// BenchmarkSonic-16        1344217               879.2 ns/op
func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonStr, _ := sonic.Marshal(c1)
		var c3 BenchmarkClass
		_ = sonic.Unmarshal(jsonStr, &c3)

	}
}

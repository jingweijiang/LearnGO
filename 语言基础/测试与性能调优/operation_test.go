package operation

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(100))
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func BenchmarkAdd(b *testing.B) {
	// 随机生成两个数进行测试
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < b.N; i++ {
		Add(r.Int31n(1000), r.Int31n(1000))
	}
}

func TestAdd2(t *testing.T) {
	type args struct {
		a int32
		b int32
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "test1", args: args{a: 1, b: 2}, want: 3},
		{name: "test2", args: args{a: 10, b: 20}, want: 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

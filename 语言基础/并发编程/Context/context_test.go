package Context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func f1(ctx context.Context) {
	ctx = context.WithValue(ctx, "key1", "value1")
	f2(ctx)
}

func f2(ctx context.Context) {
	fmt.Println(ctx.Value("key1").(string))
}

func TestContext(t *testing.T) {
	// 请求链路传值
	ctx := context.Background()
	f1(ctx)

}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("worker exit")
		return ctx.Err()
	default:
		fmt.Println("worker running")
		return nil
	}
}

func TestContext1(t *testing.T) {
	// 请求链路传值
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()

}

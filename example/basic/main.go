package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建带有取消信号的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 在 goroutine 中使用上下文
	go func(ctx context.Context) {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Operation completed")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
		}
	}(ctx)

	// 主程序等待一段时间后，取消上下文
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

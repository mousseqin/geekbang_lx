package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	// 设置一个期限
	duration := 150 * time.Millisecond

	//  创建一个既可手动取消又会发出信号的上下文。
	//	在指定的时间内取消。
	ctx , cancel := context.WithTimeout(context.Background(),duration)
	defer cancel()

	// 创建一个渠道，接收工作完成的信号。
	ch := make(chan int,1)

	// 要求goroutine为我们做一些工作。
	go func() {
		// 模拟工作
		time.Sleep(250 * time.Millisecond)

		// 报告工作已完成
		ch <- 123
	}()

	// 等待工作完成。如果时间太长，就继续前进。
	select {
	case d := <- ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

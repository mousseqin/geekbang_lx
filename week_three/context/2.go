package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 100 * time.Millisecond	// 1000Millisecond  = 1秒

func main(){
	d := time.Now().Add(shortDuration)
	ctx , cancel := context.WithDeadline(context.Background(),d)
	// 如果要实现一个超时控制，通过上面的 context 的 parent/child 机制，
	// 其实我们只需要启动一个 定时器，然后在超时的时候，直接将当前的 context 给 cancel 掉，
	// 就可以实现监听在当前和 下层的额 context.Done() 的 goroutine 的退出
	defer cancel()

	select {
	case <- time.After(1*time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Printf("err info: %+v ",ctx.Err())
	}
}

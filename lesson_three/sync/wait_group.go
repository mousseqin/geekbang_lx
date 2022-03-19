package main

import (
	"fmt"
	"sync"
)

func main() {
	res := 0
	wg := sync.WaitGroup{}
	wg.Add(10)  // 声明一个 sync.WaitGroup，然后通过 Add 方法设置计数器的值，需要跟踪多少个协程就设置多少
	for i := 0; i < 10; i++ {
		go func(val int) {
			res += val
			wg.Done() // 在每个协程执行完毕时调用 Done 方法，让计数器减 1，告诉 sync.WaitGroup 该协程已经执行完毕；
		}(i)
	}
	// 把这个注释掉你会发现，什么结果你都可能拿到
	wg.Wait() // 最后调用 Wait 方法一直等待，直到计数器值为 0，也就是所有跟踪的协程都执行完毕
	fmt.Println(res)
}

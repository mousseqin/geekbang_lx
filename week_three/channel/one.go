package main

import (
	"sync"
)
// 第一个 goroutine 在发送消息 foo 之后被阻塞，因为 还没有接收者准备好。
func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()

		//time.Sleep(time.Second * 1)
		println(`message: ` + <-c)
	}()

}

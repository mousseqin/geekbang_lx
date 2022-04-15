package main

import (
	"fmt"
	"sync"
)

// https://go.dev/blog/pipelines

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call.


	//in := gen(2, 3)
	//
	//// Distribute the sq work across two goroutines that both read from in.
	//c1 := sq(in)
	//c2 := sq(in)
	//
	//// Consume the first value from output.
	//done := make(chan struct{}, 2)
	//out := merge(done, c1, c2)
	//fmt.Println(<-out) // 4 or 9
	//
	//// Tell the remaining senders we're leaving.
	//done <- struct{}{}
	//done <- struct{}{}
	// Set up the pipeline.
	//c := gen(2, 3)
	//out := sq(c)
	//
	//// Consume the output.
	//fmt.Println(<-out) // 4
	//fmt.Println(<-out) // 9
	//
	//// Set up the pipeline and consume the output.
	//for n := range sq(sq(gen(2, 3))) {
	//	fmt.Println(n) // 16 then 81
	//}

	//in := gen(2, 3)
	//
	//// Distribute the sq work across two goroutines that both read from in.
	//c1 := sq(in)
	//c2 := sq(in)
	//
	//// Consume the merged output from c1 and c2.
	//for n := range merge(c1, c2) {
	//	fmt.Println(n) // 4 then 9, or 9 then 4
	//}
}

func merge(done <-chan struct{},cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int,1) // enough space for the unread inputs

	// 为cs中的每个输入通道启动一个输出goroutine。
	// 输出从c中复制数值到out，直到c被关闭，然后调用wg.Done。
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个goroutine，在所有的输出goroutine完成后关闭完成。 这必须在调用wg.Add后开始
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(done <-chan struct{},nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

//func gen(nums ...int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for _ , n := range nums {
//			out <- n
//		}
//		close(out)
//	}()
//	return out
//}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
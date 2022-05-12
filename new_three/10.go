package main

import (
	"fmt"
	"sync"
)

func main() {
	res := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(val int) {
			res += val
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(res)
}

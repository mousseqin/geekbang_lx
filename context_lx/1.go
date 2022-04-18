package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("one done")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("two done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all done")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

func main(){
	for routine := 1; routine <= 2 ; routine++ {
		Wait.Add(1)
		go  Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d \n",Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter // 读内存Read at 0x0000011f88a8 by goroutine 8
		time.Sleep(1*time.Nanosecond)
		value ++
		Counter = value	 // 这里赋值（写内存） Previous write at 0x0000011f88a8 by goroutine 7
	}
	Wait.Done()
}

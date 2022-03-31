package main

import (
	"sync"
)

func main(){
	c := make(chan string,2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func () {
		//defer wg.Done()
		c <- `foo`
		c <- `bar`
	}()

	go func () {
		//defer wg.Done()

		//time.Sleep(time.Second*1)
		println(`message: ` + <-c)
		println(`message: ` + <-c)
	}()
}

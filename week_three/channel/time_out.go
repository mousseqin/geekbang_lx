package main

import "time"

func main(){
	timeout := make(chan bool,1)
	go func(){
		time.Sleep(1*time.Second)
		timeout <- true
	}()

}

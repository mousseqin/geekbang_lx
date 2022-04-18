package main

import (
	"fmt"
	"time"
)

func main(){
	stop := make(chan bool)

	go func() {
		for {
			select{
			case <-stop:
				fmt.Println("stopping...")
				return
			default:
				fmt.Println("goroutine监控中..")
				time.Sleep(time.Second*2)
			}
		}
	}()

	time.Sleep(time.Second*10)
	fmt.Println("可以了，通知监控停止")

	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

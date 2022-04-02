package main

import (
	"fmt"
	"time"
)

func  main(){

	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}



	//// 准备好几个通道。
	//intChannels := [3]chan int{
	//	make(chan int, 1),
	//	make(chan int, 1),
	//	make(chan int, 1),
	//}
	//
	//// 随机选择一个通道，并向它发送元素值。
	//index := rand.Intn(3)
	//fmt.Printf("The index: %d\n", index)
	//intChannels[index] <- index
	//
	//// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	//select {
	//case <-intChannels[0]:
	//	fmt.Println("The first candidate case is selected.")
	//case <-intChannels[1]:
	//	fmt.Println("The second candidate case is selected.")
	//case elem := <-intChannels[2]:
	//	fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	//default: fmt.Println("No candidate case is selected!")
	//}

}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

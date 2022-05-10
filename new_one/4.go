package main

import "fmt"

func main(){
	arr := []int {9, 8, 7, 6}

	// 如果只需要 index 也可以去掉 写成 for index := range arr
	for index := range arr {
		fmt.Printf("only index: %d \n", index)
	}


}

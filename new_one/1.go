package main

import "fmt"

func main() {

	a1 := [3]int{7,8,9}
	fmt.Printf("a1: %v, len: %d, cap: %d", a1, len(a1), cap(a1))

	var a2 [3]int
	fmt.Printf("a2: %v, len: %d, cap: %d", a2, len(a2), cap(a2))

	// 按下标索引
	fmt.Printf("a1[1]: %d", a1[1])


}

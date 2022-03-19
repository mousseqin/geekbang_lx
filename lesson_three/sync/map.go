package main

import (
	"fmt"
	"sync"
)

func main(){
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")

	// 这里重新读取出来的，就是
	val, ok := m.Load("cat") // 取出来的是一个接口
	if ok {
		fmt.Println(val)
	}

	fmt.Println(val.(string))
	fmt.Println(len(val.(string)))
}

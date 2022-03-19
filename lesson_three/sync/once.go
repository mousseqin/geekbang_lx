package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once // 不要在函数里定义，因为会重复刷新

// 这个方法，不管调用几次，只会输出一次
func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
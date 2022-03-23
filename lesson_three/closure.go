package main

import (
	"fmt"
	"time"
)

func main() {
	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()

	fnc := ReturnClosure("Tom")()
	fmt.Println(fnc)

	Delay()
	time.Sleep(time.Second)
}

// Delay 延时绑定（非常重要 面试容易考）
func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			// 打印出得结果并不是1,2,3...递增，而都是10，因为延时绑定了
			// 闭包里面使用的闭包外的参数，其值是在最终调用的时候确定下来的
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

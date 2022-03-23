package main

import "fmt"

// main 练习断言
func main() {
	//var v interface{}
	//
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < 10; i++{
	//	v = i
	//	if (r.Intn(100) % 2) == 0 {
	//		v = "hello"
	//	}
	//	if _, ok := v.(int); ok {
	//		fmt.Printf("%v \n", v)
	//	}
	//}

	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Print(value, ",", ok)

}

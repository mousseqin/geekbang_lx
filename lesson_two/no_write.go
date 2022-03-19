package main

import "fmt"

type NwParent struct {}

func (p NwParent) Name() string {
	return "Parent"
}

func (p NwParent) SayHello() {
	fmt.Println("I am " + p.Name())
}

type NwSon struct {
	NwParent   // zuhe
}

// 定义了自己的 Name() 方法
func (s NwSon) Name() string {
	return "Son"
}

func main(){
	r := NwSon{
		NwParent{},
	}

	// 儿子没有SAYHELLO方法，也不会重写NAME
	// 最后输出：I am Parent
	// 有重写的语言会输出：I am Son
	r.SayHello()


}



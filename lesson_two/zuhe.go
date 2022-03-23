package main

import "fmt"

// Swimming 会游泳的
type swimming interface {
	Swim()
}

type duck interface {
	// 鸭子是会游泳的，所以这里组合了它
	swimming
}

type base struct {
	Name string
}

type concrete1 struct {
	base
}

func (c concrete1) SayHello() {
	// c.Name 直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.base.Name)

	// 调用了被组合的
	c.base.SayHello()

}

func (b *base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}

func main() {
	var c concrete1
	c.Name = "base_name" // 通过组合，从concrete1结构体中访问base中的NAME属性

	c.SayHello()
}

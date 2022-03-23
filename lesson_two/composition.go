package main

import "fmt"

// Swimming 会游泳的
type Swimming interface {
	Swim()
}

type Duck interface {
	// 鸭子是会游泳的，所以这里组合了它
	Swimming
}

type Base struct {
	Name string
}

type Concrete1 struct {
	Base
}

func (c *Concrete1) SayHelloZh() {
	// c.Name 直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)

	// 调用了被组合的
	c.Base.SayHello()
}

func (b *Base) SayHelloZh() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}

func main() {
	var c Concrete1
	c.Name = "test_name"

	c.SayHelloZh()

}

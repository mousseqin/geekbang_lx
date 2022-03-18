package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Color = "黑色"  // 给一个默认的颜色
	duck1.Swim()

	//duck2 := ToyDuck{}
	//duck2.Swim()

	// duck3 是 *ToyDuck
	//duck3 := new(ToyDuck)     // new 可以理解为 Go 会为你的变 量分配内存，并且把内存都置为 0
	//duck3.Swim()

	// 当你声明这样的时候，Go 就帮你分配好内存
	// 不用担心空指针的问题，以为它压根就不是指针
	//var duck4 ToyDuck
	//duck4.Swim()

	// duck5 就是一个指针了
	//var duck5 *ToyDuck
	// 这边会直接panic 掉
	//duck5.Swim()

	// 赋值，初始化按字段名字赋值
	//duck6 := ToyDuck{
	//	Color: "黄色",
	//	Price: 100,
	//}
	//duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用
	//duck7 := ToyDuck{"蓝色", 1024}
	//duck7.Swim()

	// 后面再单独赋值
	//duck8 := ToyDuck{}
	//duck8.Color = "橘色"

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint8
}

func (t *ToyDuck) Swim() {
	fmt.Printf("我是%s的鸭子，%d元一只\n", t.Color, t.Price)
}



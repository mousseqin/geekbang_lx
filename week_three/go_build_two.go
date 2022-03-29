package main

import "fmt"

type IceCreamMarker interface {

	Hello()
}

// Ben和Jerry两个结构体底层字段不一样，运行时会报：panic: runtime error: invalid memory address or nil pointer dereference
type Ben struct {
	id int
	name string
}
//type Jerry struct {
//	name *[5]byte
//	field2 int
//}
type Jerry struct {
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben Say: my name is %s \n",b.name)
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry Say: my name is %s \n",j.name)
}

func main() {
	var ben = &Ben{id:10,name:"BEN"}
	//var ben = &Ben{name:"BEN"}
	var jerry = &Jerry{name:"JERRY"}
	var maker IceCreamMarker = ben

	var loop0,loop1 func()

	loop0 = func(){
		maker = ben
		go loop1()
	}

	loop1 = func(){
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}





package main

import "fmt"

func main() {

	// fmt
	//name := "mousse"
	//age := 34
	//str := fmt.Sprint(" %s,%d \n",name,age)
	//println(str)
	//fmt.Printf("%s,%d",name,age)
	//replaceHolder()
}

func replaceHolder() {
	u := &user{
		Name: "Tom",
		Age:  17,
	}
	fmt.Printf("v => %v \n", u)   // echo     v => &{Tom 17}
	fmt.Printf("+v => %+v \n", u) // echo     &{Name:Tom Age:17}
	fmt.Printf("#v => %#v \n", u) // echo     &main.user{Name:"Tom", Age:17}
	fmt.Printf("T => %T \n", u)   // echo     *main.user
}

type user struct {
	Name string
	Age  int
}

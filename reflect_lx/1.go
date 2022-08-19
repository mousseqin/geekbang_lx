package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{"q", 18}
	t := reflect.TypeOf(u)
	fmt.Println(t) //echo main.User
	v := reflect.ValueOf(u)
	//fmt.Println(v)

	//fmt.Printf("%T\n", u)
	//fmt.Printf("%v\n", u)

	u1 := v.Interface().(User)
	fmt.Println(u1) // echo {q 18}
	t1 := v.Type()
	fmt.Println(t1) // echo main.User

	fmt.Println(t.Kind())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
}

type User struct {
	Name string
	Age  int
}

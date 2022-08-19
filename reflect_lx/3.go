package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := user{"张三", 20}
	v := reflect.ValueOf(u)
	mPrint := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))
}

type user struct {
	Name string
	Age  int
}

func (u *user) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

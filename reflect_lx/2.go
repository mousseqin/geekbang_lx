package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println(x)
}

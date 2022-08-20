package main

import (
	"github.com/davecgh/go-spew/spew"
	"unsafe"
)

func main() {
	p := Programmer{name: "future", language: "php", age: 200}

	//name := (*string)(unsafe.Pointer(&p))
	//*name = "future" // 这里已经改了

	// 已经获取到name的内存地址，就可以用偏移量来修改其它字段
	//lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	//*lang = "go"
	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "go"

	spew.Dump(p) // echo {"future","200","go"}
}

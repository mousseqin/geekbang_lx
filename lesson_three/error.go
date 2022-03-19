package main

import (
	"errors"
	"fmt"
)

type MyError struct {

}

func (e *MyError) Error() string {
	return "this is my error"
}

func main(){
	var e error =  &MyError{}
	s := e.Error()
	fmt.Println(s)

	ErrorsPkg()
}

func ErrorsPkg()  {
	err := &MyError{}
	// 使用 %w 占位符，返回的是一个新错误
	// wrappedErr 是一个新类型，fmt.wrapError
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	// 再解出来
	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped")
	}

	if errors.Is(wrappedErr, err) {
		// 虽然被包了一下，但是 Is 会逐层解除包装，判断是不是该错误
		fmt.Println("wrapped is err")
	}

	copyErr := &MyError{}
	// 这里尝试将 wrappedErr转换为 MyError
	// 注意我们使用了两次的取地址符号
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}

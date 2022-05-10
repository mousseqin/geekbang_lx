package main

import "unicode/utf8"

func main() {

	println(len("你好")) // 输出6
	println(utf8.RuneCountInString("你好")) // 输出 2
	println(utf8.RuneCountInString("你好ab")) // 输出 4

	// 反正遇到计算字符个数，比如说用户名字多长，博客多长这种字符个数
	// 记得用 utf8.RuneCountInString

	// 字符串拼接。只能发生在 string 之间
	println("Hello, " + "Go!")



}

package main

import "fmt"

// 别名，除了换了一个名字，没有任何区别
type fakeNews = News

func main() {
	var n News = fakeNews{
		Name: "hello",
	}
	n.Report()
}

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}

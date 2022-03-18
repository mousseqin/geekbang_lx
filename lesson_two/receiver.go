package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age: 10,
	}
	u.ChangeName("Tom Changed!")   // 并没有修改
	u.ChangeAge(100)				// 修改成功
	fmt.Printf("%v \n", u)			// echo {Tom 100}

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age: 12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")   // 并没有修改
	up.ChangeAge(120)				   // 修改成功

	fmt.Printf("%v \n", up)			   // echo &{Jerry 120}


	// Tip:结构体和指针之间的方法可以互相调用
}

type User struct {
	Name string
	Age int
}

// 结构体接收器
func (u User) ChangeName(newName string)  {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func main(){
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/order", order)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "这是订单")
}

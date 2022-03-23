package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there,I love %s!", r.URL.Path[1:])
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there,user%s!", r.URL.Path[1:])
}

func main() {
	//http.HandleFunc("/",handler)
	////http.HandleFunc("/",user)
	//log.Fatal(http.ListenAndServe(":8080",nil))
}

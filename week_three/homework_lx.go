package main

//import (
//	"net/http"
//
//	"github.com/micro/go-micro/server"
//)
//
//func main() {
//	srv := httpServer.NewServer(
//		server.Name("helloworld"),
//	)
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte(`hello world`))
//	})
//
//	hd := srv.NewHandler(mux)
//
//	srv.Handle(hd)
//	srv.Start()
//	srv.Register()
//}
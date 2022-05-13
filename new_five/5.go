package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Server 是http server 的顶级抽象
type Server interface {
	Routable
	// Start 启动我们的服务器
	Start(address string) error

	Shutdown(ctx context.Context) error
}

// Routable 可路由的
type Routable interface {
	// Route 设定一个路由，命中该路由的会执行handlerFunc的代码
	Route(method string, pattern string, handlerFunc handlerFunc) error
}


// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用得上
	Name    string
	handler Handler
	root    Filter
	ctxPool sync.Pool
}

func (s *sdkHttpServer) Route(method string, pattern string,
	handlerFunc handlerFunc) error {
	return s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s)
}

func (s *sdkHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	c := s.ctxPool.Get().(*Context)
	defer func() {
		s.ctxPool.Put(c)
	}()
	c.Reset(writer, request)
	s.root(c)
}

func (s *sdkHttpServer) Shutdown(ctx context.Context) error {
	// 因为我们这个简单的框架，没有什么要清理的，
	// 所以我们 sleep 一下来模拟这个过程
	fmt.Printf("%s shutdown...\n", s.Name)
	time.Sleep(time.Second)
	fmt.Printf("%s shutdown!!!\n", s.Name)
	return nil
}

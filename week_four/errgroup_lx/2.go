package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 定义 withCancel -> cancel() 方法 去取消下游的 Context
	ctx , cancel := context.WithCancel(context.Background())
	// 使用 errgroup 进行 goroutine 取消
	g , gCtx := errgroup.WithContext(ctx)
	//http server
	srv := &http.Server{Addr: ":9999"}

	// http start
	g.Go(func()error{
		return StartHttpServer(srv)
	})

	// http shutdown
	g.Go(func()error{
		<-gCtx.Done()	//阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close
		fmt.Println("http server stop")
		return srv.Shutdown(gCtx) // 关闭 http server
	})

	// signal
	sigCh := make(chan os.Signal,1)	//这里要用 buffer 为1的 chan
	signal.Notify(sigCh)

	g.Go(func()error{
		for {
			select{
			case <-gCtx.Done():	// 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return gCtx.Err()
			case si := <-sigCh:	// 因为 kill -9 或其他而终止
				cancel()
				fmt.Printf("信号信息:%+v \n ", si.String())
				return errors.New("信号错误 :" + si.String())
			}
		}
	})

	//mock error
	g.Go(func() error {
		time.Sleep(time.Second*3)
		return errors.New("I make an error")
	})

	// err group调用WAIT 自动会启用CANCEL
	if err := g.Wait(); err != nil {
		fmt.Println("end error:", err)
	}
}

//启动 HTTP server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer2)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

// 增加一个 HTTP hanlder
func HelloServer2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/da440dil/go-workgroup"
//	"net/http"
//	"time"
//)
//
//// main 基于 errgroup 实现一个 http server 的启动和关闭
////      以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
//func main() {
//	// 创建workgroup
//	var wg workgroup.Group
//	// 添加信号功能
//	wg.Add(workgroup.Signal())
//	srv := http.Server{Addr: "127.0.0.1:9999"}
//	// 增加 启动/关闭HTTP SERVER的方法
//	wg.Add(workgroup.Server(
//		// server
//		func() error {
//			fmt.Printf("serve : %v \n", srv.Addr)
//			err := srv.ListenAndServe()
//			fmt.Printf("serve close err: %v \n", err)
//			if err != http.ErrServerClosed {
//				return err
//			}
//			return nil
//		},
//		// shutdown
//		func() error {
//			fmt.Println("func shutdown")
//			ctx , cancel := context.WithTimeout(context.Background(),time.Millisecond*100)
//			defer cancel()
//
//			err := srv.Shutdown(ctx)
//			fmt.Printf("shutdown err: %v \n",err)
//			return err
//		},
//	))
//	// 5秒后取消
//	ctx , cancel := context.WithCancel(context.Background())
//	go func(){
//		time.Sleep(time.Second*5)
//		fmt.Println("ctx cancel")
//		cancel()
//	}()
//
//	wg.Add(workgroup.Context(ctx))
//
//	err := wg.Run()
//	fmt.Printf("Workgroup run stopped with error: %v \n",err)
//
//}

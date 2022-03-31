package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type httpServer struct {
	httpServer http.Server
}

func main(){
	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()

	g , _ := errgroup.WithContext(ctx)

	httpOne := NewServer(":9001")
	g.Go(func()error{
		fmt.Println("this is http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			httpOne.shutdown(context.TODO())
		}()
		return httpOne.start()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return nil
			}
		}
	})

	// inject error
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}

func NewServer(address string) *httpServer {
	return &httpServer{
		httpServer: http.Server{
			Addr:address,
		},
	}
}

func (h *httpServer) start() error {
	return h.httpServer.ListenAndServe()
}

func (h *httpServer) shutdown(ctx context.Context) error {
	return h.httpServer.Shutdown(ctx)
}




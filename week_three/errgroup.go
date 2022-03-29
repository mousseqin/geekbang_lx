package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main(){
	g := new(errgroup.Group)
	var urls  = []string {
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _ , url :=  range urls {
		//  启动一个goroutine来获取URL
		url := url
		g.Go(func()error{
			// fetch the url
			resp , err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

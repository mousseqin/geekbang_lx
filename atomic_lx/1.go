package main

import (
	"sync/atomic"
	"time"
)

func main() {
	// 保留当前的服务器配置
	var config atomic.Value
	// 创建初始配置值并存储到配置中
	config.Store(loadConfig())
	// 每10秒重新加载配置,并用新版本更新配置值。
	go func(){
		for {
			time.Sleep(time.Second*5)
			config.Store(loadConfig())
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for r:= range request() {
				c:= config.Load()
				// 使用配置c处理请求r
				_ , _ = r , c
			}
		}()
	}
}

func loadConfig() map[string]string {
	return make(map[string]string)
}

func request() chan int {
	return make(chan int)
}

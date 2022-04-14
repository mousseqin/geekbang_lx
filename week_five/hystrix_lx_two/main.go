package main

import (
	"github.com/afex/hystrix-go/hystrix"
)
func main(){
	//type CommandConfig struct {
	//	Timeout                // 等待 command 完成的时间，默认 1000（ms）
	//	MaxConcurrentRequests  // 同一个 command 支持的并发量，默认 10
	//	RequestVolumeThreshold // 触发开启熔断的最小请求量，相当于 sentinel 的静默请求数量，默认 20
	//	SleepWindow            // 熔断器打开后，控制过多久后去尝试服务是否回复，默认值是 5000（ms）
	//	ErrorPercentThreshold  // 错误百分比，当错误比例达到此值后触发熔断，默认 50（%）
	//}
	hystrix.ConfigureCommand("my_command",hystrix.CommandConfig{
		Timeout: 				1000,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  25,
	})

	// hystrix 同步
	//err := hystrix.Do("my_command", func() error {
	//	// 业务逻辑
	//	return nil
	//}, func(err error) error {
	//	// 降级逻辑
	//	return nil
	//})
	//
	//fmt.Printf("error: %v /n",err)

	// hystrix 异步
	//output := make(chan bool ,1)
	//e := hystrix.Go("my_command",func()error{
	//	// 正常业务逻辑
	//	output <- true
	//	return nil
	//},nil)
	//
	//select {
	//case out := <-output:
	//	// success
	//case err := <-e:
	//	// failure
	//}


}




















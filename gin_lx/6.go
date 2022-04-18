package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.MaxMultipartMemory = 100 << 20 这里需要注意的是保存表单缓存的内存大小，Gin默认给的是32M,通过const defaultMultipartMemory = 32 << 20 // 32 MB可以看出。如果你觉得不够，可以提前通过修改MaxMultipartMemory的值增加
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200,wechat)
	})
	r.Run(":8081")
}

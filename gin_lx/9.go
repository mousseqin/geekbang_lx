package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	a := []string{"1", "2", "3"}
	r.GET("/secureJson", func(c *gin.Context) {
		c.SecureJSON(200, a)
	})
}

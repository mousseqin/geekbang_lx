package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	// r.RedirectTrailingSlash = false   TRUE为开启重定向 
	r.GET("/users/:id",func(c *gin.Context){
		id := c.Param("id")
		c.String(200,"user_id: %s ",id)
	})
	r.Run(":8082")
}

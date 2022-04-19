package main

import "github.com/gin-gonic/gin"

func main(){
	r:= gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"message":"hello world"})
	//})
	//allUsers := []user{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
	//r.GET("/users", func(c *gin.Context) {
	//	c.IndentedJSON(200, allUsers)
	//})
	//r.GET("/users/123", func(c *gin.Context) {
	//	c.JSON(200, user{ID: 123, Name: "张三", Age: 20})
	//})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})

	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})

	r.Run(":8081")
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

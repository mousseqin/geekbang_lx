package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{{ID: 1,Name: "jack"},{ID: 2,Name: "rose"}}
	r := gin.Default()
	r.GET("/users",func(c *gin.Context){
		c.JSON(200,users)
	})

	r.Run(":8081")
}
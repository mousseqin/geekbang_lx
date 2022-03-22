package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main(){
	r:=gin.Default()

	r.GET("/list", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/adduser", createUser)
	r.Run(":8081")
}

func listUser(c *gin.Context)  {
	c.JSON(200,users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	found := false
	//类似于数据库的SQL查询
	for _, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}
	if found {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
	}
}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{ID: len(users) + 1, Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated,u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入用户名称",
		})
	}
}

var users = []User{
	{ID: 1,Name: "张三"},
	{ID: 2,Name: "李四"},
	{ID: 3,Name: "王五"},
}

//用户
type User struct {
	ID int
	Name string
}

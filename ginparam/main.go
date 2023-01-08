package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address" binding:"required"`
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(200, gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	r := gin.Default()
	// http://localhost:8082/user/save?id=1&name=xwc
	r.GET("/user/save", func(c *gin.Context) {
		var user User
		//err := c.BindQuery(&user)
		err := c.ShouldBindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		//id := c.Query("id")
		//name := c.Query("name")
		//address := c.DefaultQuery("address", "北京市")
		//c.JSON(200, gin.H{
		//	"id":   id,
		//	"name": name,
		//	"a":    address,
		//})
		c.JSON(http.StatusOK, user)
	})
	r.GET("/user/save2", func(c *gin.Context) {
		address := c.QueryArray("address")
		c.JSON(200, gin.H{
			"a": address,
		})
	})
	r.POST("/user/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})
	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}

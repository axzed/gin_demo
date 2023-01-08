package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 使用cookie存储session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		// 获取session
		session := sessions.Default(c)
		// 设置session
		if session.Get("hello") != "world" {
			fmt.Println("session is not exist")
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{
			"hello": session.Get("hello"),
		})
	})

}

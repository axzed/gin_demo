package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "hello world",
		})
	})
	//r.GET("/user/:id", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, c.Param("id"))
	//})

	v1 := r.Group("/v1")
	{
		v1.GET("/find", func(c *gin.Context) {
			c.JSON(200, "v1 find")
		})
		v1.GET("/save", func(c *gin.Context) {
			c.JSON(200, "v1 save")
		})
	}

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

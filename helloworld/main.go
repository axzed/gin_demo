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
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

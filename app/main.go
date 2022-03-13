package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run(":3000")
}

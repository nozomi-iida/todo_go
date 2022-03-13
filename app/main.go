package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nozomi-iida/todo_go/controllers"
	"github.com/nozomi-iida/todo_go/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run(":3000")
}

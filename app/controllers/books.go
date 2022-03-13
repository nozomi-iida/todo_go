package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nozomi-iida/todo_go/models"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSONP(http.StatusOK, gin.H{"data": books})
}

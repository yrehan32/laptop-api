package controllers

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, status_code int, message string) {
	c.JSON(status_code, gin.H{
		"error": true,
		"message": message,
		"data": nil,
	})
}
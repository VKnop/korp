package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	service := gin.Default()

	service.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	service.Run(":8000")
}

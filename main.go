package main

import (
	"github.com/Arueljust/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connection()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "zaya",
		})
	})
	r.Run()
}

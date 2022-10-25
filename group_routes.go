package main

import (
  "github.com/gin-gonic/gin"
)

type routes struct {
  router *gin.Engine
}

func loginEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "login, ok",
  })
}

func submitEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "submitEndpoint, ok",
  })
}

func readEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "readEndpoint, ok",
  })
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}

	router.Run(":9888")
}


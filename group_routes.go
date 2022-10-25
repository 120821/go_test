package main

import (
  "github.com/gin-gonic/gin"
)

type routes struct {
  router *gin.Engine
}

func hiEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "hi, ok",
  })
}

func haEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "ha, ok",
  })
}

func lueEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "lue, ok",
  })
}
func loginEndpoint(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(200, gin.H{
		"pathInfo": getPath,
	})
}

func submitEndpoint(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(200, gin.H{
		"pathInfo": getPath,
	})
}

func readEndpoint(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(200, gin.H{
		"pathInfo": getPath,
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
		v1.GET("/hi", hiEndpoint)
		v1.GET("/ha", haEndpoint)
		v1.GET("/lue", lueEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
		v2.GET("/hi", hiEndpoint)
		v2.GET("/ha", haEndpoint)
		v2.GET("/lue", lueEndpoint)
	}

	router.Run(":9888")
}


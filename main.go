package main

import (
  "go-test/models"
  "go-test/controllers"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  models.ConnectDatabase()

  router.POST("/products", controllers.CreateProducts)
  router.GET("/products", controllers.FindProducts)
  router.GET("/products/:id", controllers.FindProducts)
  router.PATCH("/products/:id", controllers.UpdateProducts)
  router.DELETE("/products/:id", controllers.DeleteProducts)

  router.Run("localhost:8888")
}


package controllers

import (
  "net/http"
  "samzhangjy/go-blog/models"

  "github.com/gin-gonic/gin"
)

type CreateProductInput struct {
  //Code string `json:"title" binding:"required"`
  //Price uint `json:"content" binding:"required"`
  //Code []string `form:"colors[]"`
  //Price []int `form:"prices[]"`
  Code []string `json:"code" binding:"required"`
  Price []uint `json:"price" binding:"required"`

}

func CreateProduct(c *gin.Context) {
  var input CreateProductInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  product := models.Product{Price: input.Price, Code: input.Code}
  models.db.Create(&product)

  c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProducts(c *gin.Context) {
  var products []models.Product
  models.db.Find(&products)

  c.JSON(http.StatusOK, gin.H{"data": products})
}

func FindProduct(c *gin.Context) {
  var product models.Product

  if err := models.db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": product})
}

type UpdateProductInput struct {
  Code   string `json:"code"`
  Price uint `json:"price"`
}

func UpdateProduct(c *gin.Context) {
  var product models.Product
  if err := models.db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
    return
  }

  var input UpdateProductInput

  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  updatedProduct := models.Product{Code: input.Code, Price: input.Price}

  models.db.Model(&product).Updates(&updatedProduct)
  c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
  var product models.Product
  if err := models.db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
    return
  }

  models.db.Delete(&product)
  c.JSON(http.StatusOK, gin.H{"data": "success"})
}

package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

type myForm struct {
  Colors []string `form:"colors[]"`
}

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  dsn := "host=localhost user=admin password=88888888 dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Read
  var product Product
  db.First(&product, 2) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Create
  db.Create(&Product{Code: "D77", Price: 180})
  db.Create(&Product{Code: "D77", Price: 200})
  db.Create(&Product{Code: "D99", Price: 200})
  db.Create(&Product{Code: "D99", Price: 300})

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)

  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 2)
  r := gin.Default()
  r.LoadHTMLGlob("views/*")
  r.GET("/", indexHandler)
  r.POST("/", formHandler)
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
//  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  r.Run(":8888")
}

func indexHandler(c *gin.Context) {
  c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
  var fakeForm myForm
  c.Bind(&fakeForm)
  c.JSON(200, gin.H{"color": fakeForm.Colors})
}

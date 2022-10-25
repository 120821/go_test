package main

import (
  "gorm.io/gorm"
  //  "gorm.io/driver/sqlite"
  //引入sql
  "gorm.io/driver/postgres"
  //引入 postgres
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type Apple struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(50);not null;index:ip_idx"`
	Color int    `gorm:"not null"`
	Addr string `gorm:"type:varchar(50);not null;"`
}

func main() {
  // db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  // 使用sql
  dsn := "host=localhost user=admin password=88888888 dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  // 使用postgres
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.Migrator().CreateTable(&Apple{})
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
}

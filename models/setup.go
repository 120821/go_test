package models

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
  dsn := "host=localhost user=admin password=88888888 dbname=blogs port=5432 sslmode=disable timezone=asia/shanghai"
  db, err := gorm.open(postgres.open(dsn), &gorm.config{})
  if err != nil {
    panic("failed to connect database")
  }
  //dsn := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable timezone=Asia/Shanghai"
  //database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // change the database provider if necessary

  //if err != nil {
  //  panic("Failed to connect to database!")
  //}

  database.AutoMigrate(&Product{})  // register Product model

  db = database
}

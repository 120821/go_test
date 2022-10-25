package db

import (
    "log"

    "github.com/tutorials/go/blogs/go-test/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    dbURL := "postgres://pg:88888888@localhost:5432/blogs"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})

    return db
}


package main

import (
  "fmt"
  "log"
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
   _ "github.com/lib/pq"
  "github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
  return c.SendString("Hello")
}
func postHandler(c *fiber.Ctx) error {
  return c.SendString("Hello")
}
func putHandler(c *fiber.Ctx) error {
  return c.SendString("Hello")
}
func deleteHandler(c *fiber.Ctx) error {
  return c.SendString("Hello")
}

// TODO
func main() {
  dsn := "host=localhost user=admin password=88888888 dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  app := fiber.New()

  //app.Get("/", indexHandler)

  //app.Post("/", postHandler)

  //app.Put("/update", putHandler)

  //app.Delete("/delete", deleteHandler)
  app.Get("/", func(c *fiber.Ctx) error {
    return indexHandler(c, db)
  })

  app.Post("/", func(c *fiber.Ctx) error {
    return postHandler(c, db)
  })

  app.Put("/update", func(c *fiber.Ctx) error {
    return putHandler(c, db)
  })

  app.Delete("/delete", func(c *fiber.Ctx) error {
    return deleteHandler(c, db)
  })

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

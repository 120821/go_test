package main
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//func main() {
//  //r := gin.Default()
//  //r.GET("/ping", func(c *gin.Context) {
//  //  c.JSON(http.StatusOK, gin.H{
//  //    "message": "pong",
//  //  })
//  //})
//  router := gin.Default()
//  router.GET("/hi", func(context *gin.Context) {
//    context.String(http.StatusOK, "Hello world!")
//  })
//  err := router.Run()
//  if err != nil {
//    panic("[Error] failed to start Gin server due to: " + err.Error())
//    return
//  }
//  router.Run(":9888")
//  //r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

//}
func main() {
    router := gin.Default()
    router.GET("/hi", func(context *gin.Context) {
        context.String(http.StatusOK, "Hello world!")
    })
    err := router.Run(":9888")
    if err != nil {
        panic("[Error] failed to start Gin server due to: " + err.Error())
        return
    }
}


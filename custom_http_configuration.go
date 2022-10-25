package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)
//func main() {
//  router := gin.Default()

//  s := &http.Server{
//    Addr:           ":8080",
//    Handler:        router,
//    ReadTimeout:    10 * time.Second,
//    WriteTimeout:   10 * time.Second,
//    MaxHeaderBytes: 1 << 20,
//  }
//  s.ListenAndServe()
//}

func main() {
	router := gin.Default()
	http.ListenAndServe(":9888", router)
}


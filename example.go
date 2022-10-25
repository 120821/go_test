package main
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(200, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}

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
    router.GET("/getb", GetDataB)
    router.GET("/getc", GetDataC)
    router.GET("/getd", GetDataD)
    err := router.Run(":9888")
    if err != nil {
        panic("[Error] failed to start Gin server due to: " + err.Error())
        return
    }
}


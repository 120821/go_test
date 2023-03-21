package main
import "fmt"
func main() {
  var a [5]int
  b := a
  b[3] = 7
  b[0] = 873
  b[2] = 12
  b[1] = 39
  length := len(a)
  fmt.Println(a, b) // [0 0 0 0 0] [0 0 0 7 0]
  var i int
  for i = 0; i < length;i++ {
    fmt.Println("before", b[i], "b is", b) //
    b[i] = b[i] + 2
    fmt.Println("after", b[i], "b is", b) //
  }
}

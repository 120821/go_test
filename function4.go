package main
import "fmt"
var g int
func main() {
  var a, b int
  a = 3
  b = 4
  g = a + b
  fmt.Printf("result is a %d, b = %d and g = %d\n", a, b, g)
}

package main
import "fmt"
func main() {
  for i := 2; i<15; i++{
    fmt.Println("before break i 是", i)
    break
    fmt.Println("after break i 是", i)
  }
}

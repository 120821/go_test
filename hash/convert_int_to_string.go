package main
import (
  "fmt"
  "strconv"
)
func generate_hash(number int) string{
  result := "2"
  for i := 1; i <= number; i++{
    key := strconv.Itoa(i)
    value := strconv.Itoa(i)
    fmt.Println(key)
    fmt.Println(value)
  }
  return result
}
func main() {
  fmt.Println(generate_hash(3))
}

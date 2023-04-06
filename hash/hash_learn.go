package main
import (
  "fmt"
   "strconv"
)
func generate_hash(number int) map[string]string{
  hash := make(map[string]string)
  for i := 1; i <= number; i++{
    key := "key " + strconv.Itoa(i)
    value := "string " + strconv.Itoa(i)
    hash[key] = value
    //hash["key 1"] = "string 1"
    //fmt.Println(i)
  }
  return hash
}
func main() {
  //i := generate_hash(3)
  //fmt.Println(i)
  fmt.Println(generate_hash(5))
}

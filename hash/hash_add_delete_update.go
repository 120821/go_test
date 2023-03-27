package main
import "fmt"
func main() {
  m := make(map[string]float64)

  m["pi"] = 3.14             // Add a new key-value pair
  fmt.Println("m: ", m)             // Print map: "map[pi:3.1416]"
  m["pi"] = 3.1416           // Update value
  fmt.Println("m ", m)             // Print map: "map[pi:3.1416]"

  v := m["pi"]               // Get value: v == 3.1416
  fmt.Println("v ", v)
  v = m["pie"]               // Not found: v == 0 (zero value)
  fmt.Println("v ", v)

  _, found := m["pi"]        // found == true
  fmt.Println("found ", found)
  _, found = m["pie"]        // found == false
  fmt.Println("found ", found)

  if x, found := m["pi"]; found {
    fmt.Println("x ", x)
    fmt.Println("found ", found)
  }                           // Prints "3.1416"

  delete(m, "pi")             // Delete a key-value pair
  fmt.Println("m ", m)              // Print map: "map[]"

}

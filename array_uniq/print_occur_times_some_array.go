package main
import "fmt"
func main() {
  m := make(map[int]int)
  my_array := []int {5,4,2, 3, 3, 2, 2, 1, 4, 5, 3,2,1,4}

  for i := range my_array {
    fmt.Println(i)
    fmt.Println(my_array[i])
    //if m[i] != nil {
    //  m[i] = m[i] + 1
    //} else {
    //  m[i] = 1
    //}
    switch v := m[i].(type) {
    case int:
      fmt.Println("it's an integer", v)
      m[i] = m[i] + 1
    case nil:
      fmt.Println("it's nil!")
      m[i] = 1
    }
  }

  fmt.Println(m)
}

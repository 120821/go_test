package main
import "fmt"
func main() {
  var i, j, k int
  balance := [5]float32{2.0, 1000.8, 2.3,4.2, 44.2}
  for i = 0; i < 5; i++ {
    fmt.Printf("balance[%d] = %f\n", i, balance[i])
  }
  balance2 := [...]float32{2.0, 1000.8, 2.3,4.2, 44.2}
  for j = 0; j < 5; j++ {
    fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
  }
  balance3 := [5]float32{2:4.0, 3:5.5}
  for k = 0; k < 5; k++ {
    fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
  }
}

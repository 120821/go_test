package main
import "fmt"
func main() {
  var m map[string]int                // nil map of string-int pairs
  m1 := make(map[string]float64)      // Empty map of string-float64 pairs
  m2 := make(map[string]float64, 100) // Preallocate room for 100 entries
  m3 := map[string]float64{           // Map literal
      "e":  2.71828,
      "pi": 3.1416,
  }
  fmt.Println(len(m1))
  fmt.Println(len(m))
  fmt.Println(len(m2))
  fmt.Println(len(m3))
  fmt.Println(m1)
  fmt.Println(m)
  fmt.Println(m2)
  fmt.Println(m3)
  fmt.Println(m3["e"])
  // 0
  // 0
  // 0
  // 2
  // map[]
  // map[]
  // map[]
  // map[e:2.71828 pi:3.1416]
  // 2.71828

}

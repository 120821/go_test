package main

import (
  "fmt"
)

func main() {

  vals := make([]int, 5)
  fmt.Println(vals)

  vals[0] = 12
  vals[1] = 18
  vals[2] = 13
  vals[3] = 19
  vals[4] = 38

  fmt.Println(vals)

  vals2 := []int{12, 18, 13, 19, 38}
  fmt.Println(vals2)

  vals3 := []int{}
  fmt.Println(vals3)

  // vals3[0] = 12
  // vals3[1] = 18

  vals3 = append(vals3, 1)
  vals3 = append(vals3, 2)
  vals3 = append(vals3, 5)
  vals3 = append(vals3, 6)

  fmt.Println(vals3)
}

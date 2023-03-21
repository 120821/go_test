package main

import "fmt"

func main() {

  w := make([][]string, 3)

  w1 := make([]string, 4)
  w1[0] = "war"
  w1[1] = "water"
  w1[2] = "wrath"
  w1[3] = "wrong"

  w2 := make([]string, 3)
  w2[0] = "car"
  w2[1] = "cup"
  w2[2] = "cloud"

  w3 := make([]string, 2)
  w3[0] = "boy"
  w3[1] = "brown"

  w[0] = w1
  w[1] = w2
  w[2] = w3

  fmt.Println(w)
}


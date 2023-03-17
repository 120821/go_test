package main

import (
  "fmt"
)

func main() {
  var index int = 5
  scientists := []string{
    "Einstein",
    "Turing",
    "Lovelace",
    "Curie",
    "Franklin",
    "Hodgkin",
  }

  scientists = append(scientists, "Hawkins")

  fmt.Println(scientists)
  scientists = append(scientists, "test-name")
  fmt.Println(scientists)

  elem := scientists[index]

  scientists = append(scientists[:index], scientists[index+1:]...)
  fmt.Println(scientists)
  fmt.Println(elem)

}


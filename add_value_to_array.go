package main

import (
  "fmt"
)

func main() {
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

}


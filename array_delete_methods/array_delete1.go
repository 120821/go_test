package main

import "fmt"

func remove(slice []int, s int) []int {
  return append(slice[:s], slice[s+1:]...)
}

func main() {
  beforeRemoveSlice := []int{1, 2, 3, 4, 5}
  indexRemove := 2

  fmt.Println("before remove:", beforeRemoveSlice)

  afterRemoveSlice := remove(beforeRemoveSlice, indexRemove)
  fmt.Println("after remove", afterRemoveSlice)
}

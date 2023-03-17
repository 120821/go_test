package main

import (
  "fmt"
)

func main(){

  // initialize an array
  originalArray := [4]string{"a", "b", "c", "d"}
  fmt.Println("The original array is:", originalArray)

  // initialize the index of the element to delete
  i := 2

  // check if the index is within array bounds
  if i < 0 || i >= len(originalArray) {
    fmt.Println("The given index is out of bounds.")
  } else {
    // delete an element from the array
    newLength := 0
    for index := range originalArray {
      if i != index {
        originalArray[newLength] = originalArray[index]
        newLength++
      }
    }

    // reslice the array to remove extra index
    newArray := originalArray[:newLength]
    fmt.Println("The new array is:", newArray)
  }
}

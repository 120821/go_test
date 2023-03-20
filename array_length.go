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

  fmt.Println(len(originalArray))
  // check if the index is within array bounds
  if i < 0 || i >= len(originalArray) {
    fmt.Println("The given index is out of bounds.")
  } else {
    // delete an element from the array
    newLength := 0
    for index := range originalArray {
      fmt.Println("新的长度", newLength, "i是", i, "index是", index, "第", index + 1, "次循环")
      if i != index {
        fmt.Println("赋值前=== originalArray[newLength]: ", originalArray[newLength])
        fmt.Println("赋值前=== originalArray[index]: ", originalArray[index])
        originalArray[newLength] = originalArray[index]
        fmt.Println("赋值后=== originalArray[newLength]: ", originalArray[newLength])
        fmt.Println("赋值后=== originalArray[index]: ", originalArray[index])
        newLength++
       } else {
         fmt.Println("===================== 什么都不做")
      }
    }
    // reslice the array to remove extra index
    newArray := originalArray[:newLength]
    fmt.Println("The new array is:", newArray)
    fmt.Println("The new length is:", newLength)
  }
}

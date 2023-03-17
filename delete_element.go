// Go program to find the index of
// the element from the given slice
package main

import "fmt"

func main() {

  // Declaring an slice
  numbers := []int{10, 20, 30, 40, 50, 90, 60}
  fmt.Println("Slice:", numbers)

  var index int = 3

  // Get the element at a provided index in the slice
  elem := numbers[index]
  fmt.Println("Element at index 3 is:", elem)
}


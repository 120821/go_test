package main

import "fmt"

func insertionSort(arr []int) []int {
  for i := range arr {
    preIndex := i - 1
    current := arr[i]
    for preIndex >= 0 && arr[preIndex] > current {
      arr[preIndex+1] = arr[preIndex]
      preIndex -= 1
    }
    arr[preIndex+1] = current
  }
  return arr
}

//func insertionSort(arr []int) []int {
//  for i := 0; i < len(arr); i++ {
//    for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
//      arr[j], arr[j-1] = arr[j-1], arr[j]
//    }
//  }
//  return arr
//}

func main() {
  fmt.Println(insertionSort([]int{5,3,2,1,4}))
  // prints
  // [0, 1, 2, 3, 4, 5]
}


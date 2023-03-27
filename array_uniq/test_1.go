package main
import "fmt"
func insertionSort (arr []int) []int {
  for i := 0; i<len(arr); i++{
    //fmt.Println("=====i is:", i)
   for j := i; j>0 && arr[j-1] < arr[j]; j--{
    //fmt.Println("======i is:", i)
    //fmt.Println("=====1====== arr j:", arr[j], "j is", j, "arr", arr)
    //fmt.Println("=====1====== arr j:", arr[j], "arr j-1 is", arr[j-1])
    arr[j], arr[j-1] = arr[j-1], arr[j]
    //fmt.Println("=====2====== arr j:", arr[j], "j is", j, "arr", arr)
    //fmt.Println("=====2====== arr j:", arr[j], "arr j-1 is", arr[j-1])
   }
  }
  return arr
}
func main()
  fmt.Println(insertionSort([]int {5, 4, 2, 9, 3, 4, 3, 2, 1, 4}))
  sort_array := insertionSort([]int {5, 4, 2, 9, 3, 4, 3, 2, 1, 4})
  fmt.Println(sort_array)
  length := len(sort_array)
  fmt.Println(length)
  k := length - 1
  fmt.Println(k)
  //for i := 0; i<k; i++{
  //  fmt.Println("==== before i is", i)
  //  if sort_array[i] == sort_array[i + 1]{
  //    fmt.Println("==== before", sort_array)
  //    fmt.Println("==== before sort_array[i] is", sort_array[i])
  //    fmt.Println("==== before sort_array[i-1] is", sort_array[i-1])
  //  }
  //}
}

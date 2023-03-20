package main
import "fmt"
func insertionSort (arr []int) []int {
  for i := 0; i<len(arr); i++{
    fmt.Println("=====i is:", i)
   for j := i; j>0 && arr[j-1] < arr[j]; j--{
    fmt.Println("======i is:", i)
    fmt.Println("=====1====== arr j:", arr[j], "j is", j, "arr", arr)
    fmt.Println("=====1====== arr j:", arr[j], "arr j-1 is", arr[j-1])
    arr[j], arr[j-1] = arr[j-1], arr[j]
    fmt.Println("=====2====== arr j:", arr[j], "j is", j, "arr", arr)
    fmt.Println("=====2====== arr j:", arr[j], "arr j-1 is", arr[j-1])
   }
  }
  return arr
}
func main() {
  fmt.Println(insertionSort([]int {5,3,2,1,4}))
}

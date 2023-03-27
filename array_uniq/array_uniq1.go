package main
import "fmt"
func insertionSort (arr []int) []int {
  for i := 0; i<len(arr); i++{
    for j := i; j>0 && arr[j-1] < arr[j]; j--{
      arr[j], arr[j-1] = arr[j-1], arr[j]
    }
  }
  return arr
}
func RemoveIndex(s []int, index int) []int {
  return append(s[:index], s[index+1:]...)
}

func main() {
  sort_array := insertionSort([]int {5,4,2, 3, 3, 2, 2, 1, 4, 5, 3,2,1,4})
  fmt.Println(sort_array)
  for i := 0; i<(len(sort_array)-1); i++{
    fmt.Println("i 是", i)
    for j := i; j>0; j--{
      fmt.Println("j 是", j)
      fmt.Println("sort_array[i] 是", sort_array[i])
      fmt.Println("sort_array[i+1] 是", sort_array[i+1])
      fmt.Println("sort_array[j] 是", sort_array[j])
      fmt.Println("sort_array[j+1] 是", sort_array[j+1])
      if sort_array[i+1] == sort_array[i]{
        sort_array = RemoveIndex(sort_array, i)
      }
      //arr[j], arr[j-1] = arr[j-1], arr[j]
    }
  }
  //fmt.Println(len(sort_array))
  //sort_array = RemoveIndex(sort_array, i)
  //fmt.Println(i)
  fmt.Println(sort_array)
}

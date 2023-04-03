package main
import "fmt"
func zhiShu(number int) ([]int){
  array := []int{}
  if number == 1 {
    array = append(array, 1)
  }else if number == 2{
    array = append(array, 1, 2)
  }else if number >= 2{
    array = append(array, 1, 2)
    for i := 2; i<=number; i++{
      for j := 2; j<i; j++{
        if i%j == 0 {
          break
        }else{
          if j+1 == i {
            //fmt.Println("i 是", i, "array", array)
            array = append(array, i)
          }
        }
      }
    }
  }
  return array
}

func sum(array []int) int{
  length := len(array)
  sum := 0
  for i := 1; i<=length; i++{
    fmt.Println("========== i 是", i, "array", array)
    fmt.Println("========== i%2 是", i%2)
    if i%2 == 0 {
      fmt.Println("=== 减之前，sum", sum, "在数组的第偶数位元素的符号为负", -array[i-1])
      sum = sum - array[i-1]
      fmt.Println("=== 减之后，sum", sum)
    }else{
      fmt.Println("=== 加之前, sum", sum, "在数组的第奇数位元素的符号为正", array[i-1])
      sum = sum + array[i-1]
      fmt.Println("=== 加之后, sum", sum)
    }
  }
  return sum
}

func main() {
  sort_array := zhiShu(30)
  fmt.Println(sort_array)
  sum := sum(sort_array)
  fmt.Println(sum)
}

package main
import "fmt"
func zhiShu(number int) ([]int, int){
  array := []int{}
  sum := 0
  if number == 1 {
    array = append(array, 1)
    sum = 1
  }else if number == 2{
    array = append(array, 1, 2)
    sum = 1 + 2
  }else if number >= 2{
    array = append(array, 1, 2)
    sum = 1 + 2
    for i := 2; i<=number; i++{
      for j := 2; j<i; j++{
        if i%j == 0 {
          break
        }else{
          if j+1 == i {
            //fmt.Println("i 是", i, "array", array)
            array = append(array, i)
            sum = sum + i
          }
        }
      }
    }
  }
  return array, sum
}

func main() {
  sort_array, sum := zhiShu(30)
  fmt.Println(sort_array, sum)
}

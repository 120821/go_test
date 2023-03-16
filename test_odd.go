package main
import "fmt"
func print_odd_number (number int) int{
  var n int
  for n=1; n<=number; n++{
    if n%2==0{
      //fmt.Println("结果没有余数:", n)
    }else{
      fmt.Println(n)
    }
  }
  return n
}
func main() {
  fmt.Println(print_odd_number(10))
}

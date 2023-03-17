package main
import "fmt"
import "strconv"
//import "reflect"
func print_odd_number (number int) []string{
  var n int
  //var t []string
  t := []string{}
  for n=1; n<=number; n++{
    if n%2==0{
      //fmt.Println("结果没有余数:", n)
      t = append(t, strconv.Itoa(n))
    }else{
      //fmt.Println(n)
      //fmt.Println(n, reflect.TypeOf(n))
      //t = append(t, strconv.Itoa(n))
      //fmt.Println(t, reflect.TypeOf(t))
    }
  }
  return t
}
func main() {
  fmt.Println(print_odd_number(10))
}

package main
import "fmt"
//目的: 得到n以内的阶乘
func get_factorial(number int)(result int){
  fmt.Println("===========number:", number, "result:", result)
  if number > 0{
    //阶乘公式: n!=n*(n-1)*(n-2)*(n-3)...1
    //result=n!; number=n
    result = number * get_factorial(number-1)
    fmt.Println("=======number:", number, "get_factorial n-1", get_factorial(number-1), "result1111:", result)
    return result
  }
  fmt.Println("========result:", result)
  //返回值乘1
  return 1
}
func main() {
  //调用函数
  number := 8
  fmt.Println("==============factorial:", get_factorial(int(number)))
}

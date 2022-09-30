package main
import "fmt"
func main() {
  var num1 int = 20
  var num2 int = 22
  var ret int
  ret = max(num1,num2)
  fmt.Printf("max is : %d\n", ret)
}
func max(num1,num2 int) int {
  // 声明局部变量
  var result int
  if (num1 > num2) {
    result = num1
  } else {
    result = num2
  }
  return result
}

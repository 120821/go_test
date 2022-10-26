package main
import "fmt"
func main() {
  var a int = 21
  var b int = 10
  var c int

  c = a + b
  fmt.Printf("第一行 -c 的值为 %d\n",c)
  c = a - b
  fmt.Printf("第2行 -c 的值为 %d\n",c)
  c = a * b
  fmt.Printf("第3行 -c 的值为 %d\n",c)
  c = a % b
  fmt.Printf("第4行 -c 的值为 %d\n",c)
  a++
  fmt.Printf("第5行 -a 的值为 %d\n",c)
  c = a / b
  fmt.Printf("第5行 -c 的值为 %d\n",c)
  a = 21 //重新赋值 a = 21
  a--
  fmt.Printf("第七行 -a 的值为 %d\n",a)
}

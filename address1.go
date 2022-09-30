package main
import "fmt"
func main() {
  var a int = 22
  var ip *int
  ip = &a
  fmt.Printf("a 的地址是： %x\n", &a)
  fmt.Printf("ip 的地址是： %x\n", ip)
  fmt.Printf("ip 是： %x\n", *ip)
}

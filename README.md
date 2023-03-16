# go_test
官网下载：https://go.dev/learn/
官网学习example：https://gobyexample.com/
例如： hello-world.go

```
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

$ go run hello-world.go
hello world
$ go build hello-world.go
$ ls
hello-world    hello-world.go
$ ./hello-world
hello world

命名的返回值

Go 的返回值可能会被命名。 如果是这样，它们将被视为在函数顶部定义的变量。

这些名称应用于记录返回值的含义。

A return不带参数的语句返回指定的返回值。 这被称为“裸”返回。

裸返回语句应该只用在短函数中，就像这里显示的例子一样。
vim named_results.go

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

```
$ go run named_results.go
```
7 10
```

vim test_odd.go
```
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
```
$ go run test_odd.go
```
1
3
5
7
9
11
```
vim test_add.go
```
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
```
$ go run test_add.go

```
55
```
vim test_array.go

```
package main

import "fmt"

func familyName(name string, row int, cloumn int) {
  for i:=0; i<cloumn; i++ {
    for j:=0; j<row; j++ {
      fmt.Print(name)
    }
    fmt.Println()
  }
}

func main() {
  familyName("Liam", 3, 4)
  familyName("Jenny", 14, 5)
  familyName("A", 30, 8)
  familyName("B", 3, 8)
}
```
$ go run test_array.go
return:
```
LiamLiamLiam
LiamLiamLiam
LiamLiamLiam
LiamLiamLiam
JennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJenny
JennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJenny
JennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJenny
JennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJenny
JennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJennyJenny
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
BBB
BBB
BBB
BBB
BBB
BBB
BBB
BBB

```

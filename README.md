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

vim slice.go

```
package main

import (
  "fmt"
)

func main() {

  vals := make([]int, 5)
  fmt.Println(vals)

  vals[0] = 12
  vals[1] = 18
  vals[2] = 13
  vals[3] = 19
  vals[4] = 38

  fmt.Println(vals)

  vals2 := []int{12, 18, 13, 19, 38}
  fmt.Println(vals2)

  vals3 := []int{}
  fmt.Println(vals3)

  // vals3[0] = 12
  // vals3[1] = 18

  vals3 = append(vals3, 1)
  vals3 = append(vals3, 2)
  vals3 = append(vals3, 5)
  vals3 = append(vals3, 6)

  fmt.Println(vals3)
}
```

$ go run slice.go
get result:

```
[0 0 0 0 0]
[12 18 13 19 38]
[12 18 13 19 38]
[]
[1 2 5 6]
```
vim test_array_slice.go

```
test_array_slice.go
package main
import "fmt"
func main() {
  var a [5]int
  b := a
  b[3] = 7
  b[0] = 873
  b[2] = 12
  b[1] = 39
  length := len(a)
  fmt.Println(a, b) // [0 0 0 0 0] [0 0 0 7 0]
  var i int
  for i = 0; i < length;i++ {
    fmt.Println("before", b[i], "b is", b) //
    b[i] = b[i] + 2
    fmt.Println("after", b[i], "b is", b) //
  }
}
```
$ go run test_array_slice.go

```
[0 0 0 0 0] [873 39 12 7 0]
before 873 b is [873 39 12 7 0]
after 875 b is [875 39 12 7 0]
before 39 b is [875 39 12 7 0]
after 41 b is [875 41 12 7 0]
before 12 b is [875 41 12 7 0]
after 14 b is [875 41 14 7 0]
before 7 b is [875 41 14 7 0]
after 9 b is [875 41 14 9 0]
before 0 b is [875 41 14 9 0]
after 2 b is [875 41 14 9 2]
```

vim array_height_sort.go
```
package main
import "fmt"

func sort(name []string, height []int) []string{
  for n := 0; n < len(name); n++{
    for h := n+1; h < len(name) ; h++{
      if height[n] < height[h]{
        height[n], height[h] = height[h], height[n]
        name[n], name[h] = name[h], name[n]
      }
    }
  }
  return name
}
func main() {
  var name = []string{"Mary", "John", "Emma", "hahah", "joy", "lala"}
  var height = []int{180, 165, 170, 183, 200, 179}
  fmt.Println(sort(name,height))
}

```

$ go run array_height_sort.go

```
[joy hahah Mary lala Emma John]
```

vim make_slice_of_slices.go
```
package main

import "fmt"

func main() {

  w := make([][]string, 3)

  w1 := make([]string, 4)
  w1[0] = "war"
  w1[1] = "water"
  w1[2] = "wrath"
  w1[3] = "wrong"

  w2 := make([]string, 3)
  w2[0] = "car"
  w2[1] = "cup"
  w2[2] = "cloud"

  w3 := make([]string, 2)
  w3[0] = "boy"
  w3[1] = "brown"

  w[0] = w1
  w[1] = w2
  w[2] = w3

  fmt.Println(w)
}

```
$ go run make_slice_of_slices.go
get the result:
```
[[war water wrath wrong] [car cup cloud] [boy brown]]

```

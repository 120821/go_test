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

vim make_with_map.go
```
package main

import "fmt"

func main() {

  benelux := make(map[string]string)

  benelux["be"] = "Belgium"
  benelux["nl"] = "Netherlands"
  benelux["lu"] = "Luxembourgh"

  fmt.Println(benelux)
  fmt.Printf("%q\n", benelux)
}

```

$ go run make_with_map.go
```
map[be:Belgium lu:Luxembourgh nl:Netherlands]
map["be":"Belgium" "lu":"Luxembourgh" "nl":"Netherlands"]
```
vim make_with_channel.go
```
package main

import (
  "fmt"
)

func fib(n int, c chan int) {

  x, y := 0, 1

  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func main() {

  c := make(chan int, 10)

  go fib(cap(c), c)

  for i := range c {
    fmt.Println(i)
  }
}
```

$ go run make_with_channel.go

```
0
1
1
2
3
5
8
13
21
34
```
more go tests: https://zetcode.com/all/#go
vim generate_array.go

```

```
$ go run generate_array.go
get result:
```
[[1] [1 2] [1 2 3] [1 2 3 4] [1 2 3 4 5]]
```
vim generate_array_same.go

```
package main

import (
  "fmt"
)

func generate_small_array(n int) []int{
  // 第一次循环
  result := []int{}
  for i := 1; i <= n; i++{
    result = append(result, i)
    fmt.Println("====== 每个循环1, i:", i, "result", result)
  }
  fmt.Println(result)
  return result
}

func main() {
  fmt.Println(generate_small_array(5))
  // 第二次循环
  result := [][]int{}
  for i := 1; i <= 5; i++{
    fmt.Println(generate_small_array(5))
    result = append(result, generate_small_array(i))
    fmt.Println("====== 每个循环2, i:", i, "result", result)
  }
  fmt.Println(result)
}


```
$ go run generate_array_same.go
get result:
```
[[1] [2 2] [3 3 3] [4 4 4 4] [5 5 5 5 5]]
```

### array2.go
```
package main

import "fmt"

func main() {
  // 创建二维数组
  sites := [2][2]string{}

  // 向二维数组添加元素
  sites[0][0] = "Google"
  sites[0][1] = "Runoob"
  sites[1][0] = "Taobao"
  sites[1][1] = "Weibo"

  // 显示结果
  fmt.Println(sites)
}
```
$ go run array2.go
output:
```
[[Google Runoob] [Taobao Weibo]]
```

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

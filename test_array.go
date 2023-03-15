
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


package main
import "fmt"
func sort(n []int) []int{
  fmt.Println("n:", n)
  fmt.Println("len n:", len(n))
  for i := 0; i < len(n); i++{
    fmt.Println("============i:", i)
    for j := i; j > 0 && n[j-1] > n[j]; j--{
      n[j-1], n[j] = n[j], n[j-1]
      fmt.Println("-------------j", j, n[j], n[j-1], "arr:", n)
      for m := i; m > 0 && n[m]!= n[m-1]; m--{
        fmt.Println("=========m", m, "arr:", n)
        var elem = n[m]
        fmt.Println("m", m, n[m], n[m-1], "arr:", n)
        n = append(n[:m], n[m+1:]...)
        fmt.Println("elem:", elem)
        fmt.Println("m", m, n[m], n[m-1], "arr:", n)
      }
    }
  }
  fmt.Println(n)
  return n
}
func main() {
  fmt.Println(sort([]int{1,3,3,2,3,4,2,5, 5, 4, 2, 3}))
}

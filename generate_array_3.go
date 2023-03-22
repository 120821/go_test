package main

import (
  "fmt"
)

func generate_small_array(n int, name string) []string{
  // 第一次循环
  result := []string{}
  for i := 1; i <= n; i++{
    result = append(result, name)
    fmt.Println("====== 每个循环1, i:", i, "result", result)
  }
  fmt.Println(result)
  return result
}

func main() {
  fmt.Println(generate_small_array(3, "hihi"))
  // 第二次循环
  result := [][]string{}
  for i := 1; i <= 5; i++{
    result = append(result, generate_small_array(i, "haha"))
    fmt.Println("====== 每个循环2, i:", i, "result", result)
  }
  fmt.Println(result)
}


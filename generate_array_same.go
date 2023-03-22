package main

import (
  "fmt"
)

func generate_small_array(n int) []int{
  // 第一次循环
  result := []int{}
  for i := 1; i <= n; i++{
    result = append(result, n)
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


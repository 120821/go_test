package main

import "fmt"

func main() {
  // 创建一个二维哈希
  dict := make(map[string]map[string]int)

  // 添加数据
  dict["a"] = make(map[string]int)
  dict["a"]["b"] = 1
  dict["a"]["c"] = 2
  dict["b"] = make(map[string]int)
  dict["b"]["c"] = 3

  // 访问数据
  fmt.Println(dict["a"]["b"]) // 输出 1
  fmt.Println(dict["a"]["c"]) // 输出 2
  fmt.Println(dict["b"]["c"]) // 输出 3
}

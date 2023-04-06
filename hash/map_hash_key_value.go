package main

import "fmt"

func main() {
  // 声明一个字符串型的哈希表
  hash := make(map[string]string)

  // 添加键值对
  hash["key1"] = "value1"
  hash["key2"] = "value2"

  fmt.Println(hash) // 输出结果： map[key1:value1 key2:value2]
  // 根据键获取值，并打印结果
  val1 := hash["key1"]
  fmt.Println(val1) // 输出结果：value1

  // 判断键是否存在并获取对应的值
  val2, ok := hash["key2"]
  if ok {
    fmt.Println(val2) // 输出结果：value2
  }

  // 删除键值对
  delete(hash, "key1")

  // 循环遍历哈希表
  for k, v := range hash {
    fmt.Printf("key:%s, value:%s\n", k, v)
  }
}

package main

import "fmt"

func main() {
  tempData := map[string][][]string{
    "R22":  {{"2022", "2"}, {"2023", "111"}},
    "R410": {{"2022", "1"}, {"2023", "11"}},
  }

  // 获取 R22 的数据
  r22Data := tempData["R22"]
  fmt.Println("R22 数据:")
  for _, row := range r22Data {
    fmt.Println(row[1])
  }

  // 获取 R410 的数据
  r410Data := tempData["R410"]
  fmt.Println("R410 数据:")
  for _, row := range r410Data {
    fmt.Println(row[1])
  }
}

//output:
//R22 数据:
//2
//111
//R410 数据:
//1
//11

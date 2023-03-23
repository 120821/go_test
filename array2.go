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

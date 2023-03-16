package main

import "fmt"

// 声明一个函数类型
type cb func(int) int

func main() {
  testCallBack(1, callBack)//执行函数---testCallBack
}

func testCallBack(x int, f cb) {  //定义了一个函数 testCallBack
  f(x)  //由于传进来的是callBack函数，该函数执行需要传入一个int类型参数，因此传入x
}

func callBack(x int) int {
  fmt.Printf("我是回调，x：%d\n", x)
  return x
}

package main
import "fmt"
func main() {
  list := []string {"world", "world", "hello", "hello", "hello", "hi", "hi", "hi", "hello", "hello", "world"}
  list = removeRepeatElement(list)
  fmt.Println(list)
}

func removeRepeatElement[T comparable](list []T) []T {
  // 创建一个临时map用来存储数组元素
  temp := make(map[T]struct{})
  index := 0
  // 将元素放入map中
  for _, v := range list {
    temp[v] = struct{}{}
  }

  tempList := make([]T, len(temp))
  for key := range temp {
    tempList[index] = key
    fmt.Println("key", key)
    fmt.Println("tempList", tempList)
    //fmt.Println("temp", temp)
    index++
  }
  return tempList
}
// 输出：
// [hello]


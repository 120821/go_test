package main
import (
  "fmt"
   "strconv"
)
func generate_hash(number int) map[string]map[string]string{
  hash := make(map[string]string)
  // 创建一个二维哈希
  result := make(map[string]map[string]string)
  for i := 1; i <= number; i++{
    key := "child" + strconv.Itoa(i)
    value := "value" + strconv.Itoa(i)
    hash[key] = value
    key1 := "parent" + strconv.Itoa(i)
    result[key1] = make(map[string]string)
    result[key1][key] = value
    //hash["key 1"] = "string 1"
    //fmt.Println(i)
  }
  return result
}

func main() {
  //i := generate_hash(3)
  //fmt.Println(i)
  fmt.Println(generate_hash(5))
}


//{
//  "parent1": {
//    "child1": "value1"
//  },
//  "parent2": {
//    "child2": "value2"
//  },
//  "parent3": {
//    "child3": "value3"
//  }
//}


1.使用append: 创建一个包含除要删除的元素之外的所有元素的新切片。
```
func remove(slice []int, s int) []int {
  return append(slice[:s], slice[s+1:]...)
}
```
test:
e.g.
```
package main

import "fmt"

func remove(slice []int, s int) []int {
  return append(slice[:s], slice[s+1:]...)
}

func main() {
	beforeRemoveSlice := []int{1, 2, 3, 4, 5}
	indexRemove := 2

	fmt.Println("before remove:", beforeRemoveSlice)

	afterRemoveSlice := remove(beforeRemoveSlice, indexRemove)
	fmt.Println("after remove", afterRemoveSlice)
}
```
输出：
```
before remove: [1 2 3 4 5]
after remove [1 2 4 5]
```
附上：怎么使用函数 https://www.golinuxcloud.com/golang-optional-parameters/

2.使用go的generics
```
package main

import "fmt"

func remove[T comparable](slice []T, s int) []T {
  return append(slice[:s], slice[s+1:]...)
}

func main() {
	beforeRemoveSlice := []int{1, 2, 3, 4, 5}
	strRemoveSlice := []string{"go", "linux", "cloud", "golang"}
	indexRemove := 2

	afterRemoveSlice := remove(beforeRemoveSlice, indexRemove)
	afterRemoveStrSlice := remove(strRemoveSlice, indexRemove)
	fmt.Println("after remove int slice:", afterRemoveSlice)
	fmt.Println("after remove str slice:", afterRemoveStrSlice)
}
```
输出：
```
after remove int slice: [1 2 4 5]
after remove str slice: [go linux golang]
```
3.从切片（slice）去除元素，无序
e.g.
```
package main

import "fmt"

func remove[T comparable](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}

func main() {
	beforeRemoveSlice := []int{1, 2, 3, 4, 5}
	strRemoveSlice := []string{"go", "linux", "cloud", "golang"}
	indexRemove := 2

	afterRemoveSlice := remove(beforeRemoveSlice, indexRemove)
	afterRemoveStrSlice := remove(strRemoveSlice, indexRemove)
	fmt.Println("after remove int slice:", afterRemoveSlice)
	fmt.Println("after remove str slice:", afterRemoveStrSlice)
}
```
输出：
```
after remove int slice: [1 2 5 4]
after remove str slice: [go linux golang]
```
https://www.golinuxcloud.com/golang-remove-from-slice/

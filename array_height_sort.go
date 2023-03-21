package main
import "fmt"

func sort(name []string, height []int) []string{
  for n := 0; n < len(name); n++{
    for h := n+1; h < len(name) ; h++{
      if height[n] < height[h]{
        height[n], height[h] = height[h], height[n]
        name[n], name[h] = name[h], name[n]
      }
    }
  }
  return name
}
func main() {
  var name = []string{"Mary", "John", "Emma", "hahah", "joy", "lala"}
  var height = []int{180, 165, 170, 183, 200, 179}
  fmt.Println(sort(name,height))
}

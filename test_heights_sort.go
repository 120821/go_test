package main
import "fmt"

func sortPeople(names []string, heights []int) []string {
	//ans := make([]string, len(names))
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			if heights[i] < heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
				names[i], names[j] = names[j], names[i]
			}
		}
	}
	return names
}

func main() {
  var names = []string{"Mary","John","Emma"}
  var heights = []int{180,165,170}
  fmt.Println(sortPeople(names, heights))
}

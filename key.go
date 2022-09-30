package main
import "fmt"
type Books struct {
  title string
  author string
  subject string
  book_id int
}

func main() {
  fmt.Println(Books{"go","www.baidu.com","go test",6495407})
  fmt.Println(Books{title: "go",author: "www.baidu.com",subject: "go test",book_id: 6495407})
  fmt.Println(Books{title: "go",author: "www.baidu.com"})
}

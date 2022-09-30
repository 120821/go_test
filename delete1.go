package main
import "fmt"
func main() {
  countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
  fmt.Println("原始地图")
  for country := range countryCapitalMap {
    fmt.Println(country, "首都是", countryCapitalMap [ country ])
  }
  delete(countryCapitalMap,"France")
  fmt.Println("法国被删除")
  fmt.Println("删除后地图")
  for country := range countryCapitalMap {
    fmt.Println(country, "首都是", countryCapitalMap [country])
  }
}


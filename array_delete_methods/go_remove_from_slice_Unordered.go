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

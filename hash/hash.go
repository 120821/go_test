package main

import "fmt"

func main() {
	input := "lemon"

	fmt.Printf("Hash [\"%s\"]: %d\n", input, badBadHash([]byte(input)))

	input += "s"

	fmt.Printf("Hash [\"%s\"]: %d\n", input, badBadHash([]byte(input)))
}

func badBadHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash += uint64(b)
	}

	return hash
}

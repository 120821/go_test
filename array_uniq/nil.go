package main

import (
	"fmt"
)

func main() {
	var v interface{}
	switch v := v.(type) {
	case int:
		fmt.Println("it's an integer", v)
	case nil:
		fmt.Println("it's nil!")
	}
}


package main

import (
  "fmt"
  "crypto/md5"
)

func main() {
  text := "hello world"
  hash := md5.Sum([]byte(text))
  fmt.Printf("MD5 hash of %s is %x\n", text, hash)
}

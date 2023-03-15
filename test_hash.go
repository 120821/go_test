package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Foo"

	hmd5 := md5.Sum([]byte(s))
	hsha1 := sha1.Sum([]byte(s))
	hsha2 := sha256.Sum256([]byte(s))

	fmt.Printf("   MD5: %x\n", hmd5)
	fmt.Printf("  SHA1: %x\n", hsha1)
	fmt.Printf("SHA256: %x\n", hsha2)
}

Go语言中的哈希（hash）可以用来对数据进行加密、压缩、索引等多种用途。
下面是一个简单的使用示例：

vim hash2.go
```
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
```
$ go run hash2.go
输出：
```
MD5 hash of hello world is 5eb63bbbe01eeed093cb22bb8f5acdc3
```

在以上示例中，我们使用了Go语言内置的md5包，将字符串“hello world”进行哈希处理，并将结果以十六进制形式输出。

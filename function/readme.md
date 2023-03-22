如何传递参数：
```
package main

import "fmt"

func PrintName(firstname, lastname string) {
   fullName := fmt.Sprintf("%s %s", firstname, lastname)
   fmt.Println(fullName)
}

func main() {
   firstName := "John"
   lastName := "Doe"

   PrintName(firstName, lastName)
}

```
输出：
```
John Doe
```
https://www.golinuxcloud.com/golang-optional-parameters/

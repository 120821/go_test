package main

import (
	"fmt"
	"reflect"
)

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func main() {
	calculator := Calculator{}
	methodName := "Add"
	methodArgs := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)}

	// 获取方法的反射值
	method := reflect.ValueOf(calculator).MethodByName(methodName)
	if method.IsValid() {
		// 调用方法并获取结果
		result := method.Call(methodArgs)
		if len(result) > 0 {
			// 获取方法返回值
			value := result[0]
			// 将结果转换为对应类型
			if value.Kind() == reflect.Int {
				fmt.Println("结果:", value.Int())
			}
		}
	} else {
		fmt.Println("方法不存在")
	}
}

package main

import (
	"fmt"
	"reflect"
)

func inspectVariable(variable interface{}) {
    // 獲取變量的反射值對象
    val := reflect.ValueOf(variable)

    // 獲取變量的類型
    typ := val.Type()

    // 打印類型和值
    fmt.Printf("Type: %v, Value: %v\n", typ, val)

    // 檢查變量是否為指針
    if val.Kind() == reflect.Ptr {
        // 獲取實際元素
        elem := val.Elem()

        // 檢查是否可以設置值
        if elem.CanSet() {
            // 根據類型設置新值
            if elem.Kind() == reflect.Int {
                elem.SetInt(40)
            } else if elem.Kind() == reflect.String {
                elem.SetString("Changed")
            }
            fmt.Printf("New value: %v\n", elem)
        }
    }
}

func main() {
    // 定義一些變量
    x := 10
    y := "Hello"

    // 使用反射檢查和修改它們
    inspectVariable(&x)
    inspectVariable(&y)

    // 打印修改後的值
    fmt.Println("x after modification:", x)
    fmt.Println("y after modification:", y)

}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(getType(1))
	fmt.Println(getType(false))
	fmt.Println(getType(make(chan int)))
	fmt.Println(getType("hello"))
	fmt.Println(getType([]int{}))
	fmt.Println(getType(make(map[string]int, 1)))
}

func getType(value interface{}) string {
	valueType := reflect.ValueOf(value)
	switch valueType.Kind() {
	case reflect.String:
		return "type string"
	case reflect.Int:
		return "type int"
	case reflect.Bool:
		return "type bool"
	case reflect.Chan:
		return "type chan"
	}

	return "unknown type"
}

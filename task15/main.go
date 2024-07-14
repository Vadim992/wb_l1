package main

import (
	"fmt"
	"strings"
)

func createHugeString(num int) string {
	var b strings.Builder

	for i := 0; i < num; i++ {
		b.WriteString("Р")
	}

	return b.String()
}

//old code
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
// Тянем за собой весь массив, лучше делать через copy
// Также данные могут получиться   "битыми", перевести строку v в слайс рун и потом уже перевести обрантно
//}

// new code
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	length := 100
	r := make([]rune, length)
	copy(r, []rune(v))
	justString = string(r)

	fmt.Println(justString)
}

func main() {
	someFunc()
}

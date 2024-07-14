package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "cat, cat, dog, cat, tree"
	fmt.Println(getProperSet(str))
}

func getProperSet(str string) string {
	arr := strings.Split(str, ",")
	m := make(map[string]struct{}, len(arr))

	for _, val := range arr {
		m[strings.TrimSpace(val)] = struct{}{}
	}

	var b strings.Builder
	for key := range m {
		b.WriteString(key)
		b.WriteString(", ")
	}

	res := strings.TrimSpace(b.String())

	if len(res) != 0 {
		res = res[:len(res)-1]
	}

	return res
}

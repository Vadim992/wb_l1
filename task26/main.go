package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))
}

func checkUnique(str string) bool {
	lowerCase := strings.ToLower(str)
	m := make(map[rune]struct{}, utf8.RuneCountInString(lowerCase))

	for _, char := range lowerCase {
		if _, ok := m[char]; ok {
			return false
		}

		m[char] = struct{}{}
	}

	return true
}

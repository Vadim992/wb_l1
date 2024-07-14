package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	str, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)

	out.WriteString(reverseWords(str))
}

func reverseWords(str string) string {
	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

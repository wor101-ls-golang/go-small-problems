package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	words := strings.Split(s, " ")

	for idx, word := range words {
		words[idx] = strings.Title(word)
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(CamelCase("hello case") == "HelloCase")
}

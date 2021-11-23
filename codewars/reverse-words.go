package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		reversedWord := ""

		for _, c := range word {
			char := string(c)
			reversedWord = char + reversedWord
		}
		words[i] = reversedWord
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog.") == "ehT kciuq nworb xof spmuj revo eht yzal .god")
}

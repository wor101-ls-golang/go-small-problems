package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	//	split the string into words
	// iterate over the words
	// for each word, split it into characters
	// capitalize even index characters
	// lowercase odd index characters
	// rejoin characters into a word and update words slice
	// rejoin and return words slice

	words := strings.Split(str, " ")

	for idx, word := range words {
		characters := strings.Split(word, "")

		for i, rune := range characters {
			char := string(rune)

			if i%2 == 0 {
				characters[i] = strings.ToUpper(char)
			} else {
				characters[i] = strings.ToLower(char)
			}
		}
		words[idx] = strings.Join(characters, "")
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(toWeirdCase("String"))
	fmt.Println(toWeirdCase("Weird string case"))
}

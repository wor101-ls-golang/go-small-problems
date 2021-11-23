package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToJadenCase(str string) string {
	lettersRegexp, _ := regexp.Compile("[a-zA-Z]")
	jadenString := ""
	prevChar := ""

	for i, v := range str {
		char := string(v)
		if i == 0 && lettersRegexp.Match([]byte(char)) {
			jadenString += strings.ToUpper(char)
		} else if i > 0 && lettersRegexp.Match([]byte(char)) {
			if !lettersRegexp.Match([]byte(prevChar)) {
				jadenString += strings.ToUpper(char)
			} else {
				jadenString += char
			}
		} else {
			jadenString += char
		}
		prevChar = char
	}

	return jadenString
	// iterate over string

	// // split string into a slice of words
	// words := strings.Split(str, " ")

	// // loop over the words
	// // capitalize the first letter of each word and replace the word in the slice
	// for i, word := range words {
	// 	capitalized := strings.ToUpper(string(word[0]))
	// 	if len(word) > 1 {
	// 		capitalized = capitalized + word[1:]
	// 	}

	// 	words[i] = capitalized
	// }
	// // join the slice and return it as a string
	// return strings.Join(words, " ")
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
	fmt.Println(ToJadenCase("most trees are blue") == "Most Trees Are Blue")
}

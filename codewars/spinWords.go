package main

import (
	"fmt"
	"strings"
)

func reverseWord(str string) string {
	var reversedWord string = ""

	for _, char := range str {
		reversedWord = string(char) + reversedWord
	}
	return reversedWord
}

func SpinWords(str string) string {

	// split string into slice by spaces
	strSlice := strings.Split(str, " ")

	if len(strSlice) <= 1 && len(strSlice[0]) >= 5 {
		return reverseWord(strSlice[0])
	} else if len(strSlice) <= 1 && len(strSlice[0]) < 5 {
		return strSlice[0]
	} else {
		for i, word := range strSlice {
			if len(word) >= 5 {
				strSlice[i] = reverseWord(word)
			}
		}
		return strings.Join(strSlice, " ")
	}
}

func main() {
	fmt.Println(SpinWords("azzip"))

}

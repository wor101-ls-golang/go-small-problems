package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	// create slice by splitting on '-' or '_'
	var words []string

	if strings.Contains(s, "-") {
		words = strings.Split(s, "-")
	} else {
		words = strings.Split(s, "_")
	}

	for i, v := range words {
		if i == 0 {
			continue
		}
		words[i] = strings.Title(v)
	}
	return strings.Join(words, "")

	// iterate over words
	// skip index 0
	// use title on every other value to capitalize the word
	// join and return words

}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_stealth_warrior"))
	fmt.Println(ToCamelCase("thestealthwarrior"))
}

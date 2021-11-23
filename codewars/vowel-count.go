package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := "aeiou"

	for _, v := range str {
		if strings.Index(vowels, string(v)) != -1 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra") == 5)
}

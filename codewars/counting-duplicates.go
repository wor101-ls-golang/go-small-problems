package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	duplicates := 0
	charsMap := make(map[string]int)

	for _, char := range s1 {
		charString := strings.ToUpper(string(char))
		charsMap[charString]++
	}

	for _, v := range charsMap {
		if v > 1 {
			duplicates++
		}
	}

	return duplicates
}

func main() {
	fmt.Println(duplicate_count("aabbcde") == 2)
}

package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func main() {
	fmt.Println(solution("abc", "c") == true)
	fmt.Println(solution("abcd", "c") == false)
}

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func solve(str string) string {
	lowerRegexp, _ := regexp.Compile("[a-z]")
	upperRegexp, _ := regexp.Compile("[A-Z]")

	lowerCase := lowerRegexp.FindAllString(str, -1)
	upperCase := upperRegexp.FindAllString(str, -1)

	if len(upperCase) > len(lowerCase) {
		return strings.ToUpper(str)
	}
	return strings.ToLower(str)
}

func main() {
	fmt.Println(solve("coDe") == "code")
}

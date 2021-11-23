package main

import (
	"fmt"
	"strings"
)

func Capitalization(st string) []string {
	var cap []string = []string{"", ""}

	for i, v := range st {
		letter := string(v)

		if i%2 == 0 {
			cap[0] += strings.ToUpper(letter)
			cap[1] += letter
		} else {
			cap[0] += letter
			cap[1] += strings.ToUpper(letter)
		}
	}
	return cap
}

func main() {
	fmt.Println(Capitalization("thisismystring"))

}

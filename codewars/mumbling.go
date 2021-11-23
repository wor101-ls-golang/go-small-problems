package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	// create a new slice
	// iterate over s with index
	// append capital version of s to slice followed by 0+ lowercase versions based on index
	// join slice with '-' and return

	var mySlice []string

	for i, v := range s {
		letter := string(v)
		letterString := strings.ToUpper(letter)

		for j := 0; j < i; j++ {
			letterString = letterString + strings.ToLower(letter)
		}
		mySlice = append(mySlice, letterString)
	}
	return strings.Join(mySlice, "-")
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU") == "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
}

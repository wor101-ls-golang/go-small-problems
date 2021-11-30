package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	longest := ""

	for idx, _ := range strarr {
		if idx <= len(strarr)-k {
			joined := strings.Join(strarr[idx:(idx+k)], "")
			if len(joined) > len(longest) {
				longest = joined
			}
			fmt.Println(joined)
		}

	}
	return longest
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2) == "abigailtheta")
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
}

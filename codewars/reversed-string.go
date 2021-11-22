package main

import "fmt"

func Solution(word string) string {
	reversed := ""
	for _, v := range word {
		reversed = string(v) + reversed
	}
	return reversed
}

func main() {
	fmt.Println(Solution("world") == "dlrow")
}

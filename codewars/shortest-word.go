package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	words := strings.Split(s, " ")
	shortest := len(words[0])

	for _, v := range words {
		if len(v) < shortest {
			shortest = len(v)
		}
	}
	return shortest
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps") == 3)
}

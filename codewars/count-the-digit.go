package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n, d int) int {
	var squares []string

	for i := 0; i <= n; i++ {
		squares = append(squares, strconv.Itoa(i*i))
	}
	stringOfSquares := strings.Join(squares, " ")
	fmt.Println(stringOfSquares)
	return strings.Count(stringOfSquares, strconv.Itoa(d))
}

func main() {
	fmt.Println(NbDig(25, 1) == 11)
}

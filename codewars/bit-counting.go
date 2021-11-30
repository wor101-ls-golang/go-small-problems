package main

import (
	"fmt"
	"strconv"
)

func CountBits(num uint) int {
	binary := strconv.FormatInt(int64(num), 2)
	count := 0

	for _, bit := range binary {
		if bit == 49 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(CountBits(123))
}

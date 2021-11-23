package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	charSlice := strings.Split(in, " ")
	var min int
	var minIndex int = -1
	var max int
	var maxIndex int = -1

	for i, v := range charSlice {
		value, _ := strconv.Atoi(v)

		if value < min || minIndex == -1 {
			min = value
			minIndex = i
		}

		if value > max || maxIndex == -1 {
			max = value
			maxIndex = i
		}
	}
	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4") == "42 -9")
}

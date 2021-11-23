package main

import "fmt"

func RowSumOddNumbers(n int) int {
	// row starting number equals row# * (row# -1) + 1
	currentNumber := n*(n-1) + 1
	var sum int

	for i := 0; i < n; i++ {
		sum += currentNumber
		currentNumber += 2
	}
	return sum
}

func main() {
	fmt.Println(RowSumOddNumbers(1) == 1)
	fmt.Println(RowSumOddNumbers(2) == 8)
	fmt.Println(RowSumOddNumbers(13) == 2197)
	fmt.Println(RowSumOddNumbers(19) == 6859)
	fmt.Println(RowSumOddNumbers(41) == 68921)
}

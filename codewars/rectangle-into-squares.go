package main

import "fmt"

func SquaresInRect(lng, wdth int) []int {
	// largest square will be the shortest of length or width
	// then subtract that size from largest and next is the result by the original shortest
	var originalShort int
	var originalLong int
	var largestPossibleSquare int
	var squares []int

	if lng < wdth {
		originalShort = lng
		originalLong = wdth
	} else if lng > wdth {
		originalShort = wdth
		originalLong = lng
	} else {
		return nil
	}

	largestPossibleSquare = originalShort
	squares = append(squares, largestPossibleSquare)

	if largestPossibleSquare > 1 {
		if originalShort == (originalLong - originalShort) {
			squares = append(squares, largestPossibleSquare)
		} else {
			squares = append(squares, SquaresInRect(originalShort, (originalLong-originalShort))...)
		}

	} else if originalShort == 1 {
		for x := 1; x < originalLong; x++ {
			squares = append(squares, 1)
		}
	}

	fmt.Printf("Length: %d, Width: %d = %v", lng, wdth, squares)
	fmt.Println("")

	return squares
}

func main() {
	fmt.Println(SquaresInRect(5, 3))
	fmt.Println(SquaresInRect(3, 5))
	fmt.Println(SquaresInRect(20, 14))
	fmt.Println(SquaresInRect(37, 14)) // [14, 14, 9, 5, 4, 1, 1, 1, 1]

}

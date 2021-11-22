/*
You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/

package main

import "fmt"

func PositiveSum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}) == 15)
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}) == 13)
	// Expect().To(Equal(13))
	// Expect(PositiveSum([]int{})).To(Equal(0))
	// Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
	// Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
}

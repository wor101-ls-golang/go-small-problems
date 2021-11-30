package main

import (
	"fmt"
)

func findOdd(seq []int) int {

	// create a map
	seqMap := make(map[int]int)

	fmt.Println(seq)

	// iterate over sequence
	for i := 0; i < len(seq); i++ {
		if _, exists := seqMap[seq[i]]; exists {
			seqMap[seq[i]] += 1
		} else {
			seqMap[seq[i]] = 1
		}
	}
	fmt.Println(seqMap)
	for key, value := range seqMap {
		if value%2 != 0 {
			return key
		}
	}
	return 0
}

func main() {
	fmt.Println(findOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))
	fmt.Println(findOdd([]int{5, 5, 5, 1, 1}))
	// fmt.Println(findOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))
	// fmt.Println(findOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))
	// fmt.Println(findOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))

}

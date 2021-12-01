package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	var indexes [2]int
	var answersFound bool

	for i := 0; i < len(numbers); i++ {
		if answersFound {
			break
		}
		for j := (i + 1); j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				indexes[0], indexes[1] = i, j
				answersFound = true
				break
			}
		}
	}

	return indexes
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
}

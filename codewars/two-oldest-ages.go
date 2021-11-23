package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) (oldest [2]int) {
	sort.Ints(ages)
	oldest[0], oldest[1] = ages[len(ages)-2], ages[len(ages)-1]
	return oldest
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}

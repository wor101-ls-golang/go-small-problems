package main

import (
	"fmt"
	"math"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func MxDifLg(a1, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	sort.Sort(byLength(a1))
	sort.Sort(byLength(a2))

	// fmt.Println("Array1: ", a1[0])
	// fmt.Println("Array2: ", a2[0])
	dif := len(a1[len(a1)-1]) - len(a2[0])
	shortestArray1 := int(math.Abs(float64(dif)))

	dif = len(a1[0]) - len(a2[len(a2)-1])
	shortestArray2 := int(math.Abs(float64(dif)))

	if shortestArray1 < shortestArray2 {
		return shortestArray2
	}
	return shortestArray1
}

func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}

	fmt.Println(MxDifLg(s1, s2))

}

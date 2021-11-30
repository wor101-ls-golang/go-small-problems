package main

import (
	"fmt"
	"sort"
)

type bySize []float32

func (arr bySize) Len() int {
	return len(arr)
}

func (arr bySize) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr bySize) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func FindUniq(arr []float32) float32 {
	sort.Sort(bySize(arr))
	if arr[0] != arr[1] {
		return arr[0]
	}
	return arr[len(arr)-1]
}

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 1, 2, 1, 1}) == 2)
}

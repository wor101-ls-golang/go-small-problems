package main

import "fmt"

func MaxMultiple(d, b int) int {
	var max int

	for i := b; i > 0; i-- {
		if i%d == 0 {
			max = i
			break
		}
	}
	return max
}

func main() {
	fmt.Println(MaxMultiple(2, 7) == 6)
}

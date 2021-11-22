package main

import "fmt"

func Sumation(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(Sumation(1) == 1)
	fmt.Println(Sumation(8) == 36)
}

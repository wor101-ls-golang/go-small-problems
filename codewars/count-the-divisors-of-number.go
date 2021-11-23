package main

import "fmt"

func Divisors(n int) int {
	var count int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(Divisors(4) == 3)
	fmt.Println(Divisors(5) == 2)
	fmt.Println(Divisors(30) == 8)
}

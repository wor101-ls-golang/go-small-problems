package main

import "fmt"

func sum(slice []float64) (sum float64) {
	for _, v := range slice {
		sum += v
	}
	return
}

func Tribonacci(signature [3]float64, n int) []float64 {
	sequence := signature[0:]

	for x := 3; x < n; x++ {
		nextNum := sum(sequence[(x - 3):x])
		sequence = append(sequence, nextNum)
	}

	fmt.Println(sequence)
	return sequence[0:n]
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
}

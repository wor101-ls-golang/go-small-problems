package main

import "fmt"

func calcVolume(n int) int {
	volume := n * n * n

	if n > 1 {
		volume += calcVolume(n - 1)
	}
	return volume
}

func FindNb(m int) int {
	answer := -1

	for n := 1; n < m; n++ {
		volume := calcVolume(n)
		if volume == m {
			answer = n
		} else if volume > m {
			break
		}
	}
	return answer
}

func main() {
	fmt.Println(FindNb(1071225) == 45)
}

package main

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v",
		numbers[0], numbers[1], numbers[2], numbers[3], numbers[4],
		numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

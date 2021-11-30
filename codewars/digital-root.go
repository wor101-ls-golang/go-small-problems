package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {
	stringInt := strconv.Itoa(n)
	stringSlice := strings.Split(stringInt, "")
	sum := 0

	for _, val := range stringSlice {
		intVal, _ := strconv.Atoi(val)
		sum += intVal
	}

	if sum > 9 {
		sum = DigitalRoot(sum)
	}
	return sum
}

func main() {
	fmt.Println(DigitalRoot(456))

}

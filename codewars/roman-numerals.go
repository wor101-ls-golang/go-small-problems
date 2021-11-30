package main

import "fmt"

var numerals = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func Decode(roman string) int {

	// iterate over each character in roman
	// check the value of the current character to see if it is greater or less than the value of the next character
	// if it is greater or the same, add it to the final total
	// if it is less, then subtract from the next characters value. add to total and skip next character
	var total int
	var usedNext bool = false

	for idx, val := range roman {
		if usedNext {
			usedNext = false
			continue
		}

		currentNumeral := string(val)
		currentValue := numerals[currentNumeral]
		if idx < (len(roman) - 1) {
			nextNumeral := string(roman[idx+1])
			nextValue := numerals[nextNumeral]

			if currentValue >= nextValue {
				total += currentValue
			} else {
				total += (nextValue - currentValue)
				usedNext = true
			}
		} else {
			total += currentValue
		}
	}

	return total
}

func main() {
	fmt.Println(Decode("MCMXC") == 1990)
	fmt.Println(Decode("MCMXC"))
}

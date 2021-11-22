/*
Introduction
The first century spans from the year 1 up to and including the year 100, the second century - from the year 101 up to and including the year 200, etc.

Task
Given a year, return the century it is in.

Examples
1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20
*/

package main

import "fmt"

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return (year / 100) + 1
}

func main() {
	fmt.Println(century(int(1990)) == 20)
	fmt.Println(century(int(1705)) == 18)
	fmt.Println(century(int(1900)) == 19)
	fmt.Println(century(int(1601)) == 17)
	fmt.Println(century(int(2000)) == 20)

}

// Expect(century(int(1990))).To(Equal(20))
// Expect(century(int(1705))).To(Equal(18))
// Expect(century(int(1900))).To(Equal(19))
// Expect(century(int(1601))).To(Equal(17))
// Expect(century(int(2000))).To(Equal(20))
// Expect(century(int(89))).To(Equal(1))

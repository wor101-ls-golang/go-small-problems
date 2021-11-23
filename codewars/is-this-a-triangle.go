package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	// a + b > c
	// b + c > a
	// c + a > b
	return (a+b) > c && (b+c) > a && (c+a) > b
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2) == false)
	fmt.Println(IsTriangle(4, 2, 3) == true)
}

package main

import "fmt"

func Arithmetic(a, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "divide":
		if b == 0 {
			panic("Cannot divide by zero")
		}
		return a / b
	case "multiply":
		return a * b
	default:
		panic("Unknown operator")
	}

}

func main() {
	fmt.Println(Arithmetic(5, 2, "add") == 7)
}

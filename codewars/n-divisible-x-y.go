package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	if n%x == 0 && n%y == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsDivisible(3, 3, 4) == false)
	fmt.Println(IsDivisible(12, 3, 4) == true)

}

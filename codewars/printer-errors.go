package main

import (
	"fmt"
)

func PrinterError(s string) string {
	var errors int
	length := len(s)

	for _, v := range s {
		if string(v) > "m" {
			errors += 1
		}
	}

	return fmt.Sprintf("%d/%d", errors, length)
}

func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

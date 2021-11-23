package main

import "fmt"

func DNAStrand(dna string) string {
	comp := ""

	for _, v := range dna {
		sym := string(v)

		switch sym {
		case "A":
			comp += "T"
		case "T":
			comp += "A"
		case "G":
			comp += "C"
		case "C":
			comp += "G"
		}
	}
	return comp
}

func main() {
	fmt.Println(DNAStrand("AAAA") == "TTTT")
	fmt.Println(DNAStrand("ATTGC") == "TAACG")
}

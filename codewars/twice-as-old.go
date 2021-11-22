package main

import "fmt"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	if dadYearsOld > sonYearsOld*2 {
		return dadYearsOld - (sonYearsOld * 2)
	} else {
		return (sonYearsOld * 2) - dadYearsOld
	}
}

func main() {

	fmt.Println(TwiceAsOld(36, 7) == 22)
	fmt.Println(TwiceAsOld(55, 30) == 5)

	// Expect( TwiceAsOld(36,7)).To(Equal(22))
	// Expect( TwiceAsOld(55,30)).To(Equal(5))
	// Expect( TwiceAsOld(42,21)).To(Equal(0))
	// Expect( TwiceAsOld(22,1)).To(Equal(20))
	// Expect( TwiceAsOld(29,0)).To(Equal(29))
}

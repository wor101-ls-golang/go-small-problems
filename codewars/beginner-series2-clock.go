package main

import "fmt"

func Past(h, m, s int) int {
	milliseconds := 0

	milliseconds += h * 60 * 60
	milliseconds += m * 60
	milliseconds += s
	milliseconds *= 1000

	return milliseconds
}

func main() {
	fmt.Println(Past(0, 1, 1) == 61000)
	// It("Past(0, 1, 1)", func() { Expect(Past(0, 1, 1)).To(Equal(61000)) })
	// It("Past(1, 1, 1)", func() { Expect(Past(1, 1, 1)).To(Equal(3661000)) })
	// It("Past(0, 0, 0)", func() { Expect(Past(0, 0, 0)).To(Equal(0)) })
	// It("Past(1, 0, 1)", func() { Expect(Past(1, 0, 1)).To(Equal(3601000)) })
	// It("Past(1, 0, 0)", func() { Expect(Past(1, 0, 0)).To(Equal(3600000)) })

}

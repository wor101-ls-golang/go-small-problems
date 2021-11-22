package main

import "fmt"

func OtherAngle(a int, b int) int {
	return 180 - a - b
}

func main() {
	fmt.Println(OtherAngle(30, 60) == 90)
	fmt.Println(OtherAngle(60, 60) == 60)
	fmt.Println(OtherAngle(43, 78) == 59)
	fmt.Println(OtherAngle(10, 20) == 150)

	// It("OtherAngle(30, 60)", func() { Expect(OtherAngle(30, 60)).To(Equal(90)) })
	// It("OtherAngle(60, 60)", func() { Expect(OtherAngle(60, 60)).To(Equal(60)) })
	// It("OtherAngle(43, 78)", func() { Expect(OtherAngle(43, 78)).To(Equal(59)) })
	// It("OtherAngle(10, 20)", func() { Expect(OtherAngle(10, 20)).To(Equal(150)) })
}

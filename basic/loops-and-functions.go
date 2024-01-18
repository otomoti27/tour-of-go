package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		newZ := z - (z*z-x)/(2*z)
		fmt.Println(i, ":", newZ)
		if math.Abs(newZ-z) < 0.000000000000001 {
			break
		}
		z = newZ
	}
	return z
}

func main() {
	fmt.Println("scratch  :", Sqrt(2))
	fmt.Println("math.Sqrt:", math.Sqrt(2))
	fmt.Println("scratch  :", Sqrt(3))
	fmt.Println("math.Sqrt:", math.Sqrt(3))
	fmt.Println("scratch  :", Sqrt(4))
	fmt.Println("math.Sqrt:", math.Sqrt(4))
}

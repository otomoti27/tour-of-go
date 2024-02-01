package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		newZ := z - (z*z-x)/(2*z)
		if math.Abs(newZ-z) < 0.000000000000001 {
			break
		}
		z = newZ
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

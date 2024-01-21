package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 関数も値である
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// 関数を引数にとる関数
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// クロージャ
	x, y := 0, 0
	return func() int {
		if y == 0 {
			x, y = 0, 1
		} else {
			x, y = y, x+y
		}
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

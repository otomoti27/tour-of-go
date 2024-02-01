package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() { // T implements I
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() { // F implements I
	fmt.Println(f)
}

func main() {
	var i I

	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// i = F(math.Pi)
	// describe(i)
	// i.M()

	var t *T
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

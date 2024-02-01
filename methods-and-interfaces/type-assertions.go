package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // type assertion
	fmt.Println(s)

	s, ok := i.(string) // type assertion with error handling
	fmt.Println(s, ok)

	f, ok := i.(float64) // type assertion with error handling
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

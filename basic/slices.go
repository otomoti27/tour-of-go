package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slices()
	// SlicesAreLikeReferencesToArrays()
	// SliceLiteral()
	// SliceDefaults()
	// SliceLengthAndCapacity()
	// NilSlices()
	// CreatingASliceWithMake()
	SliceOfSlices()
}

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 4は含まれない（半開区間）
	fmt.Println(s)            // [3 5 7]
}

func SlicesAreLikeReferencesToArrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	// Sliceの要素を変更すると、元の配列の要素も変更される（参照を保持しているイメージ）
	b[0] = "XXX"
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}

func SliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q) // [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{13, true},
	}
	fmt.Println(s) // [{2 true} {3 true} {5 true} {7 true} {11 true} {13 true}]
}

func SliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3 5 7]
	fmt.Println(s)

	s = s[:2] // [3 5]
	fmt.Println(s)

	s = s[1:] // [5]
	fmt.Println(s)
}

func SliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func NilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
	}
}

func CreatingASliceWithMake() {
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)

	c := b[:2] // len(c)=2, cap(c)=5
	printSlice2("c", c)

	d := c[2:5] // len(d)=3, cap(d)=3
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func SliceOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"}, // 0行目
		[]string{"_", "_", "_"}, // 1行目
		[]string{"_", "_", "_"}, // 2行目
	}

	// 以下のように書くこともできる
	// board := [][]string{
	// 	{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// }

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// X _ X
	// O _ X
	// _ _ O
}

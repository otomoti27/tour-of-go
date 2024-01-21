package main

import "fmt"

func adder() func(int) int {
	// 関数を返す関数
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 関数を返す関数
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		// クロージャ
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

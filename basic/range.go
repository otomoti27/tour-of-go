package main

import "fmt"

func main() {
	// Range()
	RangeContinued()
}

func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { // rangeはスライスやマップを反復処理する際に使う
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func RangeContinued() {
	pow := make([]int, 10) // makeでスライスを作成する
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}
	for _, value := range pow { // インデックスは不要な場合は_に代入する
		fmt.Printf("%d\n", value)
	}
}

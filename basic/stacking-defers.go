package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 遅延された関数呼び出しはスタックに積まれます。関数が返るとき、その関数に対する遅延呼び出しはLIFO(後入れ先出し)の順序で実行されます。
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

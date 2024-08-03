package main

import "fmt"

func fibonacci() func(int) int {
	a, b := 0, 1
	return func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fibonacci() // f は fibonacci() の返り値である関数
	for i := 0; i < 10; i++ {
		// ループしても fibonacci() の中の a, b は保持される
		// そのため、引数 i に基づいて異なる累積結果を保持する
		// 結果として、i 番目のフィボナッチ数を出力する
		fmt.Println(f(i))
	}
}

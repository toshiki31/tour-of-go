package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if ステートメントは、 for ループと同様に、条件の前に短いステートメントを書くことができます。
	// このステートメントで宣言された変数は、 if のスコープ内だけで有効です。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

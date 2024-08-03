package main

import "fmt"

// Go の関数はクロージャ（closure）です。
// クロージャは、それ自身の外部から変数を参照する関数値です。

// adder 関数は、整数 sum を初期化し、無形関数 func(int) int を返す
// この無名関数は adder 関数のスコープにある sum 変数にアクセスできる -> クロージャー
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // i が 無名関数の引数 x に渡される
			neg(-2*i), // -2*i が 無名関数の引数 x に渡される
		)
	}
}
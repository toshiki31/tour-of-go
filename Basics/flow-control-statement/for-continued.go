package main

import "fmt"

func main() {
	sum := 1
	// 初期化(i := 1)と後処理(i++)の記述はなくてもよい
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	var a [2]string // a は 長さ2のstring の配列
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // 配列をまとめて宣言するときは末尾でオブジェクト内で宣言
	fmt.Println(primes)
}

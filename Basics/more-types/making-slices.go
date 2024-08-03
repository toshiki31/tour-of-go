package main

import "fmt"

func main() {
	// スライスは組み込みの make 関数を使用して作成することができる
	a := make([]int, 5) // 長さ 5 のスライスを作成し、ゼロ値で初期化する
	printSlice("a", a) // a: [0 0 0 0 0]

	b := make([]int, 0, 5) // 長さ 0 で容量 5 のスライスを作成する
	printSlice("b", b) // b: []

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

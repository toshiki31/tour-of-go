package main

import "fmt"

func main() {
	var s []int   // スライスのデフォルト値（ゼロ値）は nil
	printSlice(s) // []

	// nil スライスに要素を追加する
	s = append(s, 0)
	printSlice(s) // [0]

	// スライスを拡張する
	s = append(s, 1)
	printSlice(s) // [0 1]

	// 複数要素を追加する
	s = append(s, 2, 3, 4)
	printSlice(s) // [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

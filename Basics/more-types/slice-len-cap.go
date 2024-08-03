package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 長さ0のスライスを作成
	s = s[:0]
	printSlice(s)

	// スライスを拡張
	s = s[:4]
	printSlice(s)

	// スライスを縮小
	s = s[2:] // 元のスライスの2番目の要素からスライスを作成 -> 元のスライスから2つ減る
	printSlice(s)

	// スライスを元の容量を超えてかくちょう
	// s = s[:7] // 容量を超えてスライスを拡張するとランタイムエラーが発生する
	// printSlice(s)
}

func printSlice(s []int) {
	// len()；スライスの長さ
	// cap()；スライスの容量
	// スライスの容量は、スライスの先頭柄元の配列の最後までの要素数を指す
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

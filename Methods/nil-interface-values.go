package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // (<nil>, <nil>)
	i.M() // 呼び出す具体的なメソッドを示す型がないため、ランタイムエラー
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
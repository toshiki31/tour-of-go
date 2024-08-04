package main

import "fmt"

func main() {
	var i interface{} // 空のインターフェース（任意の型を保持できる）
	describe(i) // (<nil>, <nil>)

	i = 42 
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
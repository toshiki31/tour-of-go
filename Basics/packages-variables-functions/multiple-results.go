package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// := は変数宣言と代入を同時に行う
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

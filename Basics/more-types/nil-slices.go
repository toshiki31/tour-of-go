package main

import "fmt"

func main() {
	var s []int // スライスのデフォルト値（ゼロ値）は nil
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
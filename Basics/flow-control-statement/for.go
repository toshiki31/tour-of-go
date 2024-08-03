package main

import "fmt"

func main() {
	sum := 0
	// for の条件式は括弧()が不要
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	pow := make([]int, 10) // 10個の要素を持つスライスを作成
	fmt.Println(pow)       // [0 0 0 0 0 0 0 0 0 0]
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// index や value は "_" で代入を捨てることができる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

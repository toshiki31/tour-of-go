package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	// インターフェースの値と具体的な値を表示
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"Hello"} // T型のポインタをインターフェース型の変数に代入
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

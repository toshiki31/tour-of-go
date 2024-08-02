package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1 // uint64型は、64ビットの符号なし整数を保持できる
	z      complex128 = cmplx.Sqrt(-5 + 12i) // complex128型は、float64型の実数部と虚数部を持つ
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// 特別な理由がない限り、整数の変数が必要な場合は int を使う
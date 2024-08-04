package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs() はメソッドと呼ばれる
// メソッド：関数と同じように定義されたものが、特定の型に関連付けられる
// (v Vertex) はレシーバ引数と呼ばれる
// このメソッド Abs() が Vertex 型の変数 v に関連付けられる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

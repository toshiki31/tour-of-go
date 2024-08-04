package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// メソッドは、特定の型(MyFloat)に関連付けられる関数
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2) // math.Sqrt2 は平方根の近似値(1.4142135623730951)
	fmt.Println(f.Abs())
}

// レシーバー(f MyFloat) を伴うメソッド(Abs())の宣言は、
// レシーバー型が同じパッケージ（この場合は main パッケージ）にある必要がある
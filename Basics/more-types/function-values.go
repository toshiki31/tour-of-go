package main

import (
	"fmt"
	"math"
)

// 関数も他の変数のように引数として渡すことができる
func compute(fn func(float64, float64) float64) float64 {
	// 戻り値として関数を返すこともできる
	return fn(3, 4)
}

func main () {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	// math.Pow は第一引数を第二引数で累乗する関数
	fmt.Println(compute(math.Pow))
}
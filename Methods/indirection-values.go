package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbcFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbcFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // p.Abs() は (*p).Abs() として解釈される
	fmt.Println(AbcFunc(*p))
}

// メソッドが変数レシーバの場合でも、
// ポインタレシーバの場合と同様に、変数とポインタの両方で呼び出すことができる
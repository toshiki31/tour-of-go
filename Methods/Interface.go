package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	var b Abser
	// var c Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	b = &v

	// c = v // これはエラーになる（Abs メソッドが Vertex ではなく *Vertex に定義されているため）

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

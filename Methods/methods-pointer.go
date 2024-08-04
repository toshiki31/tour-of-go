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

func (s Vertex) Abs2() float64 {
	return math.Sqrt(s.X*s.X + s.Y*s.Y)
}

// ポインタレシーバ( *Vertex )を持つメソッド(Scale)は、レシーバが指す変数を変更できる
// ポインタレシーバはレシーバ自身を変更できる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 変数レシーバだと、変数のコピーを操作するだけなので、元の変数は変更されない
func (s Vertex) Scale2(f float64) {
	s.X = s.X * f
	s.Y = s.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println("ポインタレシーバ：",v.Abs()) // 50
	s := Vertex{3, 4}
	s.Scale2(10)
	fmt.Println("変数レシーバ",s.Abs2()) // 5
}
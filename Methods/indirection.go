package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // 利便性のため、v.Scale(2) は (&v).Scale(2) として解釈される
	ScaleFunc(&v, 10)
	// ScaleFunc(v, 10) // コンパイルエラー

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v,p)
}

// メソッドがポインタレシーバである場合、
// 呼び出し時に変数、またはポインタのいずれかをレシーバとして取ることができる
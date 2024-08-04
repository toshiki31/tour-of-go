package main

import (
	"fmt"
	"image"
)

// image パッケージ
// type Image interface {
// 	ColorModel() color.Model
// 	Bounds() Rectangle
// 	At(x, y int) color.Color
// }

// Bounds() は Rectangle という構造体を返す
// Rectangle は MinというPoint型のフィールドと MaxというPoint型のフィールドを持つ
// type Rectangle struct {
// 	Min, Max Point
// }

// At(x, y int) color.Color は color.Color を返す
// color.Color は interface で、RGBA() というメソッドを持つ
// type Color interface {
// 	RGBA() (r, g, b, a uint32)
// }

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

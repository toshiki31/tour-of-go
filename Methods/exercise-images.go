package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	// カスタムカラーの計算
	v := uint8((x^y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}

// pic.ShowImage はimage.Imageインターフェースを実装した型を引数に取る
// この型は、ColorModel() color.Model, Bounds() image.Rectangle,
// At(x, y int) color.Colorの3つのメソッドを持つ
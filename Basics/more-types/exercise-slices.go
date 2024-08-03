package main

import (
	"golang.org/x/tour/pic"
)

// Pic 関数を実装する
// この関数は、長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを
// 割り当てたものを返すように実装する必要がある

func Pic(dy, dx int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}
func main() {
	// pic.Show 関数は func(int, int) [][]uint8 型の関数を引数に取る
	// Pic の引数は Show 関数の内部で与えられる
	pic.Show(Pic)
}

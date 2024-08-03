package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 が暗黙的に入る
	v3 = Vertex{}      // X:0 と Y:0 が暗黙的に入る
	p  = &Vertex{1, 2} // Vertex の　アドレスを取得
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

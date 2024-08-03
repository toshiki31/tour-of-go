package main

import "fmt"

// 構造体(struct)は、フィールド(field)の集まり
type Vertex struct {
	X int // フィールド(field)
	Y int // フィールド(field)
}

func main() {
	fmt.Println(Vertex{1, 2})
}

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // ドット(.)を使ってフィールドにアクセスできる
	fmt.Println(v.X)
}

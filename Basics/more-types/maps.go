package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map はキーと値を関連づける
// ゼロ値は nil で、nil の map はキーを持っていない
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

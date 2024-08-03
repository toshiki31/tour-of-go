package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v               // 構造体のアドレスを取得
	fmt.Println(p)        // &{1 2} -> ポインタが指す値を表示（デフォルト）
	fmt.Printf("%p\n", p) // 0xc00000e030 -> ポインタのアドレスを直接表示
	p.X = 1e9             // フィールド X にアクセスするのに、(*p).X と書く代わりに、p.X と書ける
	fmt.Println(v)
}

// 1e9 は 1 * 10^9 を表す

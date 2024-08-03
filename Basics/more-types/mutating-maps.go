package main

import "fmt"

func main() {
	m := make(map[string]int)

	// m への要素の挿入・更新
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// m への要素の削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかは、2つ目の値(ok)で確認できる
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// もし elem や ok を初めて宣言するのであれば、
// 以下のように短い宣言を使うことができる
// v, ok := m[key]
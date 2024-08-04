package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// インターフェース i が具体的な型 string を保持し、 その値を s に代入することを試みる
	s := i.(string) // i が持ってるであろう string 型の値を s　に代入(s = "hello")
	fmt.Println(s)

	s, ok := i.(string) // i が string 型の値を持っているかどうかを確認する
	fmt.Println(s, ok)  // s = "hello", ok = true

	f, ok := i.(float64) // i が float64 型の値を持っているかどうかを確認する
	fmt.Println(f, ok)   // f = 0, ok = false

	f = i.(float64) // i は float64 型の値を持っていないので、panic が発生する
	fmt.Println(f)
}

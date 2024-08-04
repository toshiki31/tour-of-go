package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// fmt パッケージに定義されている Stringerを使用
// type Stringer interface {
// 	String() string
// }

// このメソッドは Stringer インターフェースの String メソッドを実装している
// このメソッドを実行することで、fmt パッケージの関数 fmt.Println が呼び出された際に、
// Person 型をこのメソッドで定義した文字列に変換して出力する
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	b := "hello"
	fmt.Println(a, z)
	fmt.Println(b)
}

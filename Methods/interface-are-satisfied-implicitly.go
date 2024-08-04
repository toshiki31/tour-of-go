package main

import "fmt"

// M メソッドを含むインターフェース
// どんな型でもM()メソッドを持っていれば、その型はIインターフェースを実装していると見なされます。
type I interface {
	M()
}

type T struct {
	S string
}

// T型は明示的にIインターフェースを実装していると宣言する必要はありません。
// T型がM()メソッドを持っているので、暗黙的にIインターフェースを実装しています。

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// T型の変数を作成し、Iインターフェースを通してその変数を操作する
	var i I = T{"hello"}
	i.M()
}
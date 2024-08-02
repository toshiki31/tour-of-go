package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // var 宣言を省略して、型を省略して変数を宣言できる
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// 関数の外では、キーワードで始まる宣言( var, func, など)が必要で、 := での短い変数宣言は使用できない

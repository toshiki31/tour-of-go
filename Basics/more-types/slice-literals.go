package main 

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false,true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

// リテラル：ソースコード内で値を直接表現するために使用される記述方法
// 配列リテラルは固定長の配列を初期化する
// ex) [3]bool{true, true, false}

// スライスのリテラルは、長さのない配列リテラルのようなものです
// ex) []bool{true, true, false}
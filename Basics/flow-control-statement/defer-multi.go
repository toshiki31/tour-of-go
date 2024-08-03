package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// deferへ渡した関数が複数ある場合、その呼び出しはスタックされる
// スタックされた呼び出しは、LIFO(last-in-first-out)の順番で実行される（FILOと同義語）
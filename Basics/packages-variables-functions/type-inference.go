package main

import "fmt"

func main() {
	v := 42 // 明示的な型を指定しない場合、変数の型は右側の変数から推論される
	fmt.Printf("v is of type %T\n", v)
}

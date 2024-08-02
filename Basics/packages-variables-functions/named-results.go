package main

import "fmt"

// 6行目で関数の戻り値に名前をつけることができる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // (x, y int)で返り値の名前を指定しているので return x, y と書かなくも良い
}

func main() {
	fmt.Println(split(17))
}

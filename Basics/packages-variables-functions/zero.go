package main

import "fmt"

func main() {
	var i int     // デフォルト値は0
	var f float64 // デフォルト値は0
	var b bool    // デフォルト値はfalse
	var s string  // デフォルト値は""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

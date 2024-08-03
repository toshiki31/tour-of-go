package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	preZ := z
	count := 0
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("z: %v\n", z)
		fmt.Printf("|preZ-z|: %v\n", math.Abs(preZ-z))
		// 差の絶対値が0.0001未満になったら終了
		if math.Abs(preZ-z) < 0.0001 {
			// 10回未満で収束する場合は、"under 10"を表示
			if count < 10 {
				fmt.Println("under 10")
			}else{
				fmt.Println("over 10")
			}
			break
		}
		preZ = z
		count++
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(2),
		Sqrt(100),
	)
}

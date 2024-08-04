package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	preZ := z
	count := 0
	for {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("z: %v\n", z)
		// fmt.Printf("|preZ-z|: %v\n", math.Abs(preZ-z))
		// 差の絶対値が0.0001未満になったら終了
		if math.Abs(preZ-z) < 0.0001 {
			// 10回未満で収束する場合は、"under 10"を表示
			if count < 10 {
				fmt.Println("under 10")
			} else {
				fmt.Println("over 10")
			}
			break
		}
		preZ = z
		count++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

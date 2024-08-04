package main

import "fmt"

// byte型: 8ビット符号なし整数(uint8)
// [4]byte型: 4つの要素を持つ配列
type IPAddr [4]byte

// IPAddr型にString()メソッドを実装する
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

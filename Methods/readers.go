package main

import(
	"fmt"
	"io"
	"strings"
)

// io パッケージは、データストリームを読むことを表現する io.Reader インターフェースを規定している
// io.Reader インターフェースは、Read メソッドを持つ
// func (T) Read(b []byte) (n int, err error)
// Read は、データを与えられたバイトスライスに読み込み、読み込んだバイトの長さとエラーの情報を返す

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
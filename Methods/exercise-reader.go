package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// Read は、与えられたバイトスライスに文字 'A' を埋め込み、
// そのバイトスライスの長さと nil を返す
func (r MyReader) Read(b []byte) (int, error) {
	for i  := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

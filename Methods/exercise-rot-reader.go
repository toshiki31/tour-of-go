package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		// A-Z は 65-90
		// a-z は 97-122
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = (b[i]-'A'+13)%26 + 'A'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = (b[i]-'a'+13)%26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

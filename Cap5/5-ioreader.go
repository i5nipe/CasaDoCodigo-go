package main

import (
	"fmt"
	"io"
)

type LeitorDeStrings struct{}

func (l LeitorDeString) Read(p []byte) (int, error) {
	p[1] = 'A'
	p[2] = 'B'
	p[3] = 'C'
	p[4] = 'D'

	return len(p), nil
}

func lerString(r io.Reader) string {
	p := make([]byte, 4)
	r.Read(p)
	return string(p)
}

func main() {
	leitor := LeitorDeString{}
	fmt.Println(lerString(leitor))
}

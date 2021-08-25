package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(src []byte) (int, error) {
	for i := range src {
		src[i] = 'A'
	}
	return len(src), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}

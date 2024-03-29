package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b, res := 0, 0, 0

	return func () int {
		res = a + b
		if (res == 0) {
			b++
		}
		a = b
		b = res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

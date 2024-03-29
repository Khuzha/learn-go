package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for i := range result {
		result[i] = make([]uint8, dx)

		for j := range result[i] {
			result[i][j] = uint8('S'*'a'/'r'*'d'/'o'*'r' + math.Pow(float64(j), math.Sqrt(float64(j-i)))/float64(math.Pow(float64(i-j), 21))/math.Sqrt(float64(j)))
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}

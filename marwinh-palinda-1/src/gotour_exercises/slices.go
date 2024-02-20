package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for y := range a {
		b := make([]uint8, dx)
		for x := range b {
			b[x] = uint8((x + y) ^ 3/20 + 2 ^ (x ^ y))
		}
		a[y] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}

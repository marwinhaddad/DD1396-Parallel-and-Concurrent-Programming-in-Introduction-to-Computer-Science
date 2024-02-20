package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	zn, z, diff, n := 0.0, x, 1.0, 1
	for diff > math.Pow(10, -15) {
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(zn - z)
		zn = z
		fmt.Println(n, z)
		n++
	}
	return z
}

func main() {
	a := 3.0
	gosqrt := math.Sqrt(a)
	mysqrt := Sqrt(a)
	fmt.Println("math.Sqrt: ", gosqrt, "Sqrt: ", mysqrt)
	fmt.Println("Diff: ", math.Abs(gosqrt-mysqrt))
}

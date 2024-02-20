package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	a, b := 0, 1
	fmt.Println(a)
	return func() int {
		a, b = b, a+b
		return a
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 12; i++ {
		fmt.Println(f())
	}
}

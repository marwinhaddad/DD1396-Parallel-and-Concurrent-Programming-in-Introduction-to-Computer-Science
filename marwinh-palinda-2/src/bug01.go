package main

import (
	"fmt"
)

// I want this program to print "Hello world!", but it doesn't work.

// Fix 1
func main() {
	ch := make(chan string)
	go func() { ch <- "Hello world!" }()
	fmt.Println(<-ch) // <- receives value from channel
}

// Fix 2
//func main() {
//	ch := make(chan string, 1)
//	ch <- "Hello world!"
//	fmt.Println(<-ch)
//}

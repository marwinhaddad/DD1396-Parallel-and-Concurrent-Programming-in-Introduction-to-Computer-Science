package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for t := range time.Tick(delay) {
		fmt.Printf("The time is %s: %s\n", t.Format("15.04.05"), text)
	}
}

func main() {
	// Remind("Hello, World!", 2*time.Second)
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}

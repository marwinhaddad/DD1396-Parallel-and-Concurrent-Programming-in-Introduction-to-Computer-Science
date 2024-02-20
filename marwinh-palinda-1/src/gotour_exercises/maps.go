package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	stringSlice := strings.Fields(s)
	for _, key := range stringSlice {
		_, ok := m[key]
		if ok {
			m[key]++
		} else {
			m[key] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

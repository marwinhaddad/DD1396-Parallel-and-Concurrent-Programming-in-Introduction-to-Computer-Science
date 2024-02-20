package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

const (
	DataFile = "loremipsum.txt"
	Routines = 16
)

func WordCount(text string) map[string]int {
	total := make(chan map[string]int)
	slices := make(chan []string, Routines)
	freqs := make(map[string]int)

	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	arr := strings.Fields(strings.ToLower(text))

	for i := 0; i < Routines; i++ {
		go SliceText(slices, arr)
		go WordCountSlice(slices, total)
	}
	for i := 0; i < Routines; i++ {
		tempMap := <-total
		for key, value := range tempMap {
			freqs[key] += value
		}
	}
	return freqs
}

func SliceText(slices chan []string, arr []string) {
	size := (len(arr) + Routines - 1) / Routines
	for a := 0; a < len(arr); a += size {
		b := a + size
		if b > len(arr) {
			b = len(arr)
		}
		slices <- arr[a:b]
	}
}

func WordCountSlice(slices chan []string, ch chan map[string]int) {
	m := make(map[string]int)
	for _, word := range <-slices {
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	ch <- m
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	data, _ := ioutil.ReadFile(DataFile)

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

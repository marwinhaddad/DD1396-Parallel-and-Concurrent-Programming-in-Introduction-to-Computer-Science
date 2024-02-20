package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

const DataFile = "loremipsum.txt"

// WordCount Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	arr := strings.Fields(strings.ToLower(text))
	for _, word := range arr {
		_, ok := freqs[word]
		if ok {
			freqs[word] += 1
		} else {
			freqs[word] = 1
		}
	}
	return freqs
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
	// read in DataFile as a string called data
	data, _ := ioutil.ReadFile(DataFile)

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

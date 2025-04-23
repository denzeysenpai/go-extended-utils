package performance

import (
	"fmt"
	"time"
)

type Status int16

// ? BENCHMARKING AND ALGO

type Algorithm struct {
	code    string
	details string
}
type Benchmark_info struct {
	start      time.Time
	duration   int64
	details    string
	algorithms []Algorithm
}

func (benchmark_info *Benchmark_info) Add_new_algo(code string, details string) {
	benchmark_info.algorithms = append(benchmark_info.algorithms, Algorithm{code, details})
}

func NewBenchmark(details string) *Benchmark_info {
	start := time.Now()
	return &Benchmark_info{
		start:      start,
		duration:   0,
		details:    details,
		algorithms: []Algorithm{},
	}
}

func (benchmark_info *Benchmark_info) Stop() {
	benchmark_info.duration = time.Since(benchmark_info.start).Milliseconds()
	color := "\033[32m"
	if benchmark_info.duration > 600 {
		color = "\033[33m"
	} else if benchmark_info.duration > 900 {
		color = "\033[31m"
	}
	var algorithms string
	for index, algo := range benchmark_info.algorithms {
		algorithms = algorithms + fmt.Sprintf("\n\t[%d] Code: %s %s", index+1, algo.code, algo.details)
	}
	final_string := fmt.Sprintf("\nDuration: %s%dms\033[0m\nDetails: %s\nAlgo: %s\n", color, benchmark_info.duration, benchmark_info.details, algorithms)
	fmt.Print(final_string)
}

func (benchmark_info *Benchmark_info) Current_algo() Algorithm {
	if len(benchmark_info.algorithms) > 0 {
		return benchmark_info.algorithms[len(benchmark_info.algorithms)-1]
	} else {
		return Algorithm{
			code:    "No algo found",
			details: "Something went wrong in the implementation",
		}
	}
}

// ? BENCHMARKING AND ALGO

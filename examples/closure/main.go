package main

import (
	"fmt"
	"github.com/didip/stopwatch"
)

func main() {
	a := 1
	f := func() {
		for i := 1; i <= 10; i++ {
			a = a + 1
		}
	}

	latency := stopwatch.Measure(f)

	fmt.Printf("Latency in nanoseconds: %v, Result: %v\n", latency, a)
}

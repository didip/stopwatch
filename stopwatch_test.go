package stopwatch

import (
	"testing"
)

func TestSimpleWithClosureScope(t *testing.T) {
	a := 1
	f := func() {
		for i := 1; i <= 10; i++ {
			a = a + 1
		}
	}

	latency := Measure(f)

	if latency < 0 {
		t.Errorf("Fail to measure latency of a simple function. Latency: %v, Result: %v", latency, a)
	}
}

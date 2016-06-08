package stopwatch

import (
	"net/http"
	"time"
)

// Measure latency in nanoseconds
func Measure(f func()) int64 {
	start := time.Now().UnixNano()
	f()
	return time.Now().UnixNano() - start
}

// LatencyHandler is a middleware that measures latency given http.Handler struct.
func LatencyHandler(resultChan chan int64, next http.Handler) http.Handler {
	middle := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixNano()
		next.ServeHTTP(w, r)
		resultChan <- (time.Now().UnixNano() - start)
	}

	return http.HandlerFunc(middle)
}

// LatencyFuncHandler is a middleware that measures latency given request handler function.
func LatencyFuncHandler(resultChan chan int64, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return LatencyHandler(resultChan, http.HandlerFunc(nextFunc))
}

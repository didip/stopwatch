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
func LatencyHandler(resultChan chan int64, methods []string, next http.Handler) http.Handler {
	middle := func(w http.ResponseWriter, r *http.Request) {
		foundMethod := false
		for _, method := range methods {
			if r.Method == method {
				foundMethod = true
				break
			}
		}

		if foundMethod {
			start := time.Now().UnixNano()
			next.ServeHTTP(w, r)
			resultChan <- (time.Now().UnixNano() - start)
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(middle)
}

// LatencyFuncHandler is a middleware that measures latency given request handler function.
func LatencyFuncHandler(resultChan chan int64, methods []string, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return LatencyHandler(resultChan, methods, http.HandlerFunc(nextFunc))
}

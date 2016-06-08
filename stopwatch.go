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

func New() *Stopwatch {
	s := &Stopwatch{}
	s.ResultChan = make(chan int64)
	return s
}

// Stopwatch struct for capturing latency result asynchronously via channel.
type Stopwatch struct {
	ResultChan chan int64
}

// LatencyHandler is a middleware that measures latency given http.Handler struct.
func LatencyHandler(s *Stopwatch, next http.Handler) http.Handler {
	middle := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixNano()
		next.ServeHTTP(w, r)
		s.ResultChan <- (time.Now().UnixNano() - start)
	}

	return http.HandlerFunc(middle)
}

// LatencyFuncHandler is a middleware that measures latency given request handler function.
func LatencyFuncHandler(s *Stopwatch, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return LatencyHandler(s, http.HandlerFunc(nextFunc))
}

package main

import (
	"fmt"
	"github.com/didip/stopwatch"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// 1. Create a channel to receive latency result
	helloHandlerLatencyChan := make(chan int64)

	// 2. Pull latency result asynchronously.
	go func() {
		for {
			select {
			case latency := <-helloHandlerLatencyChan:
				fmt.Printf("Latency of HelloHandler in nanoseconds: %v\n", latency)
			}
		}
	}()

	fmt.Println("Starting HTTP server on :12345")
	http.Handle("/", stopwatch.LatencyFuncHandler(helloHandlerLatencyChan, []string{"GET"}, HelloHandler))
	http.ListenAndServe(":12345", nil)
}

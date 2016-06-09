package main

import (
	"bufio"
	"fmt"
	"github.com/didip/stopwatch"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	graphiteAddrString := "localhost:2003"
	hostname, _ := os.Hostname()
	hostnameUnderscore := strings.Replace(hostname, ".", "_", -1)
	hostnameUnderscore = strings.Replace(hostnameUnderscore, "-", "_", -1)

	// 1. Create TCP connection
	graphiteAddr, err := net.ResolveTCPAddr("tcp", graphiteAddrString)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, graphiteAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)

	// 2. Create a channel to receive latency result
	helloHandlerLatencyChan := make(chan int64)

	// 3. Pull latency result asynchronously.
	go func() {
		for {
			for latency := range helloHandlerLatencyChan {
				payload := fmt.Sprintf("graphite-reporting.%s.requests.HelloHandler %d %d\n", hostnameUnderscore, latency, time.Now().Unix())
				fmt.Printf("Payload for graphite: %v", payload)

				fmt.Fprintf(writer, payload)
				writer.Flush()
			}
		}
	}()

	fmt.Println("Starting HTTP server on :12345")
	http.Handle("/", stopwatch.LatencyFuncHandler(helloHandlerLatencyChan, []string{"GET"}, HelloHandler))
	http.ListenAndServe(":12345", nil)
}

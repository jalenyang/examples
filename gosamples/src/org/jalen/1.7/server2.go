package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int64

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler2(resp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Printf("URL.path = %v\n", req.URL.Path)
}

func counter(resp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(resp, "count=%d", count)
	mu.Unlock()
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func sleep(n int, ch chan int) {
	time.Sleep(time.Duration(n) * 1e9)
	ch <- n
	close(ch)
}

func requestHandle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[:1]
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}

	numString := r.URL.Query().Get("n")
	n, err := strconv.Atoi(numString)
	if err != nil {
		n = 1
	}

	ch := make(chan int)
	go sleep(n, ch)
	nRes := <-ch
	fmt.Fprintf(w, "sleep in %d sec\n", nRes)
}

func main() {
	http.HandleFunc("/", requestHandle)
	http.ListenAndServe(":8124", nil)
}
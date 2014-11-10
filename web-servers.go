package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	
}

func (h Hello) ServerHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServer("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
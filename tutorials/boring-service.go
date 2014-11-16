package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) <-chan string {	// Returns receive-only channel of strings.
	c := make(chan string)
	go func() {	// We launch the goroutine from inside the function.
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // Return the channel to the caller.
}

func main() {
	joe := boring("Joe")	// Function returning a channel.
	ann := boring("Ann")
	for i:=0; i<5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You're boring; I'm leaving.")
}
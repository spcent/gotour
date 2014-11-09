package main

import "fmt"

func fibonacci() func() int {
	first := 1
	second := 1
	return func() int {
		second = first + second
		first = second - first
		
		return second
	}
}

func main() {
	f := fibonacci()
	for i:=0; i<10; i++ {
		fmt.Println(f())
	}
}
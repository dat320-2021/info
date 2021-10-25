package main

import (
	"fmt"
	"sync"
)

var counter int

func updateCounter(arg string, wg *sync.WaitGroup) {
	fmt.Printf("%s: begin\n", arg)
	for i := 0; i < 1e7; i++ {
		counter = counter + 1
	}
	fmt.Printf("%s: end\n", arg)
	wg.Done()
}

func main() {
	fmt.Printf("main: begin (counter = %d)\n", counter)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go updateCounter("A", wg)
	go updateCounter("B", wg)
	wg.Wait()
	fmt.Printf("main: done with both goroutines ( last counter = %d )\n", counter)
}

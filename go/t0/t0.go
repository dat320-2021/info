package main

import (
	"fmt"
	"sync"
)

// Converting Fig 26.2 to Go
// To compile: go build
// To run: ./t0

func myThread(arg string, wg *sync.WaitGroup) {
	fmt.Printf("%s\n", arg)
	wg.Done()
}

func main() {
	fmt.Println("main: begin")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// pthread_create(..."A")
	go myThread("A", wg)
	go myThread("B", wg)
	// pthread_join(): wait for given thread to complete
	wg.Wait()
	fmt.Println("main: end")
}

package main

import (
	"fmt"
	"info/go/semaphore"
	"sync"
)

// Code herein corresponds to Fig 31.12 Adding Mutual Exclusion Correctly.

const MAX = 10

var (
	buffer    [MAX]int
	fillIndex int
	useIndex  int
	mutex     = semaphore.Semaphore{Value: 1} // binary semaphore == a lock
	empty     = semaphore.Semaphore{Value: MAX}
	full      = semaphore.Semaphore{Value: 0}
)

func put(val int) {
	buffer[fillIndex] = val           // Line F1: T1 gets descheduled here
	fillIndex = (fillIndex + 1) % MAX // Line F2
}

func get() int {
	tmp := buffer[useIndex]         // Line G1
	useIndex = (useIndex + 1) % MAX // Line G2
	return tmp
}

func producer(itemsToProduce int) {
	for i := 0; i < itemsToProduce; i++ {
		empty.Wait() // p1
		mutex.Wait() // p0 Lock
		put(i)       // p2
		mutex.Post() // p4 Unlock
		full.Post()  // p3: signal to consumer that an item has been added to the buffer
	}
	wg.Done()
}

func consumer(name int) {
	for {
		full.Wait()   // c1
		mutex.Wait()  // c0 Lock
		item := get() // c2
		mutex.Post()  // c4 Unlock
		empty.Post()  // c3: signal to producer that item in buffer has been consumed
		fmt.Printf("C%d: %d\n", name, item)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go producer(10)
	go producer(10)
	go consumer(1)
	go consumer(2)
	wg.Wait()
}

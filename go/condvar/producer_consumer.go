package main

import (
	"fmt"
	"sync"
)

var (
	buffer int
	count  int // the condition to be checked
	mutex  sync.Mutex
	cv     = sync.NewCond(&mutex)
)

func put(val int) {
	assert(count == 0)
	count = 1
	buffer = val
}

func get() int {
	assert(count == 1)
	count = 0
	return buffer
}

func producer(itemsToProduce int) {
	for i := 0; i < itemsToProduce; i++ {
		mutex.Lock()    // p1
		if count == 1 { // p2
			cv.Wait() // p3
		} // p4
		put(i)      // p5
		cv.Signal() // p6: signal to consumer that an item has been added to the buffer
		mutex.Unlock()
	}
	wg.Done()
}

func consumer() {
	for {
		mutex.Lock()     // c1
		for count == 0 { // c2
			cv.Wait() // c3
		}
		item := get()  // c4
		cv.Signal()    // c5: signal to producer that item in buffer has been consumed
		mutex.Unlock() // c6
		fmt.Println(item)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go producer(10)
	go consumer()
	// go consumer()
	wg.Wait()
}

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

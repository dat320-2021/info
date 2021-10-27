package main

import (
	"fmt"
	"sync"
)

const MAX = 10

var (
	buffer    [MAX]int
	count     int // the condition to be checked
	fillIndex int
	useIndex  int
	mutex     sync.Mutex
	empty     = sync.NewCond(&mutex)
	fill      = sync.NewCond(&mutex)
)

// invariant: count is in the range [0, MAX)

func put(val int) {
	buffer[fillIndex] = val
	fillIndex = (fillIndex + 1) % MAX
	count++
}

func get() int {
	tmp := buffer[useIndex]
	useIndex = (useIndex + 1) % MAX
	count--
	return tmp
}

func producer(itemsToProduce int) {
	for i := 0; i < itemsToProduce; i++ {
		mutex.Lock()       // p1
		for count == MAX { // p2
			empty.Wait() // p3
		} // p4
		put(i)        // p5
		fill.Signal() // p6: signal to consumer that an item has been added to the buffer
		mutex.Unlock()
	}
	wg.Done()
}

func consumer(name int) {
	for {
		mutex.Lock()     // c1
		for count == 0 { // c2
			fill.Wait() // c3
		}
		item := get()  // c4
		empty.Signal() // c5: signal to producer that item in buffer has been consumed
		mutex.Unlock() // c6
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

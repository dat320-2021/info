package main

import (
	"fmt"
	"sync"
)

var (
	// done = 0 - not finished
	// done = 1 - finished
	done int
	lock sync.Mutex
	cond = sync.NewCond(&lock)
)

func child() {
	fmt.Println("child")
	lock.Lock()
	defer lock.Unlock()
	done = 1 // child is finished
	cond.Signal()
}

func thread_join() {
	lock.Lock()
	defer lock.Unlock()
	for done != 1 {
		// wait will release the lock
		// put itself to sleep
		// and re-acquire it when signaled
		cond.Wait()
	}
}

func main2() {
	fmt.Println("parent: begin")
	go child()
	// wait for child to finish
	thread_join()
	fmt.Println("parent: end")
}

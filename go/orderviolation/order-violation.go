package main

import (
	"fmt"
	"sync"
	"time"
)

// mThread is shared state
var mThread *ThreadInfo

var (
	// synchronization variables
	mutex *sync.Mutex
	cond  *sync.Cond
)

type ThreadInfo struct {
	State int
}

// Thread 1:
func main() {
	mutex = &sync.Mutex{}
	cond = sync.NewCond(mutex)
	mThread = createThread(setup)
	mutex.Lock()
	cond.Signal()
	mutex.Unlock()
	fmt.Println("Hello")
	time.Sleep(time.Second)
}

func createThread(m func()) *ThreadInfo {
	go m()
	fmt.Println("DAT520")
	return &ThreadInfo{State: 101}
}

// Thread 2:
func setup() {
	mutex.Lock()
	for mThread == nil {
		cond.Wait()
	}
	mutex.Unlock()
	fmt.Println("DAT320", mThread.State)
}

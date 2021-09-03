package main

import (
	"fmt"
	"os"
	"sync"
)

var x = 0

func inc(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func incMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// default func f is without protection
	f := func() { inc(&wg) }
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-m":
			fmt.Print("mutex: ")
			var m sync.Mutex
			f = func() { incMutex(&wg, &m) }
		case "-c":
			fmt.Print("channel: ")
			ch := make(chan bool, 1)
			f = func() { incChannel(&wg, ch) }
		}
	} else {
		fmt.Print("no protection: ")
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}

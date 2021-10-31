package main

import (
	"log"
	"sync"
	"time"
)

type SemaphoredWaitGroup struct {
	sem chan bool
	wg  sync.WaitGroup
}

func (s *SemaphoredWaitGroup) Add(delta int) {
	s.wg.Add(delta)
	s.sem <- true
}

func (s *SemaphoredWaitGroup) Done() {
	<-s.sem
	s.wg.Done()
}

func (s *SemaphoredWaitGroup) Wait() {
	s.wg.Wait()
}

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

func worker(id int, wg WaitGroup) {
	// defer wg.Done()
	// defer log.Printf("#%d done", id)
	log.Printf("#%d starting", id)
	time.Sleep(time.Second)
	log.Printf("#%d done", id)
	wg.Done()
}

func main() {
	wg := SemaphoredWaitGroup{sem: make(chan bool, 5)}
	// var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	log.Printf("all done")
}

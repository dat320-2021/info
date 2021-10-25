package counter

import (
	"fmt"
	"sync"
	"testing"
)

// To run the tests:
//   go test -v -cpu 1,2,4,6,12
//   go test -count 10
//
// To run the benchmarks:
//   go test -v -run none -bench . -cpu 1,2,4,6,12
//   go test -v -run none -bench . -benchtime 10s
//   go test -v -run none -bench . -benchtime 100x
//
// Benchstat tool:
//   go install golang.org/x/perf/cmd/...@latest
//   # run the benchmarks before and after making changes
//   go test -v -run none -bench . -cpu 1,2,4,6,12 > old.txt
//   # make changes
//   go test -v -run none -bench . -cpu 1,2,4,6,12 > new.txt
//   benchstat old.txt new.txt

const (
	totalCount = 1_000_000
	numCPUs    = 6
)

func TestCounter(t *testing.T) {
	const expected = 1
	c := NewCounter()
	c.increment()
	if c.get() != expected {
		t.Errorf("after increment counter = %d, expected %d", c.get(), expected)
	}
}

func TestConcurrentEncapsulatedCounter(t *testing.T) {
	c := NewCounter()

	var wg sync.WaitGroup
	// mu := sync.Mutex{}
	for i := 0; i < totalCount; i++ {
		wg.Add(1)
		go func(j int) {
			// mu.Lock()
			v := c.get()
			// interrupted
			v++
			c.setVal(v)
			// mu.Unlock()
			// c.setVal(c.get() + 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	if c.get() != totalCount {
		t.Errorf("after increment counter = %d, expected %d", c.get(), totalCount)
	}
}

func TestConcurrentCounter(t *testing.T) {
	expected := totalCount
	c := NewCounter()

	var wg sync.WaitGroup
	for i := 0; i < totalCount; i++ {
		wg.Add(1)
		go func(j int) {
			c.increment()
			wg.Done()
		}(i)
	}
	res := 0
	mu := sync.Mutex{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			k := c.get()
			mu.Lock()
			res = k // res is critical section
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(res)

	got := c.get()
	if got != expected {
		t.Errorf("after increment counter = %d, expected %d", got, expected)
	}
}

func BenchmarkNoLockCounter(b *testing.B) {
	// baseline: no goroutines or locks
	tc := totalCount * numCPUs
	// b.N can be tweaked with the -benchtime argument, e.g., 100x or 10s
	counter := NewCounter()
	for i := 0; i < b.N; i++ {
		for j := 0; j < tc; j++ {
			counter.noLockInc()
		}
	}
}

func BenchmarkConcurrentNoLockCounter(b *testing.B) {
	// another baseline: goroutines, but no locks
	counter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(numCPUs)
		for k := 0; k < numCPUs; k++ {
			go func() {
				for j := 0; j < totalCount; j++ {
					counter.noLockInc()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentCounter(b *testing.B) {
	// concurrent counter with goroutines and locks
	counter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(numCPUs)
		for k := 0; k < numCPUs; k++ {
			go func() {
				for j := 0; j < totalCount; j++ {
					counter.increment()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

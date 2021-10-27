package bbq_test

import (
	"info/go/bbq"
	"sync"
	"testing"
)

func TestBBQInsertAndRemove(t *testing.T) {
	q := bbq.New()
	var wg sync.WaitGroup
	insert := func() {
		for i := 0; i < 10; i++ {
			q.Insert(i)
		}
		wg.Done()
	}
	remove := func() {
		for i := 0; i < 10; i++ {
			_ = q.Remove()
		}
		wg.Done()
	}
	const n = 10
	wg.Add(2 * n)
	// start 2n goroutines:
	// - n insert goroutines
	// - n remove goroutines
	for i := 0; i < n; i++ {
		go insert()
		go remove()
	}
	// wait for 2n goroutines to finish
	wg.Wait()
}

func TestChannelQueueInsertAndRemove(t *testing.T) {
	q := bbq.NewChannelQueue()
	var wg sync.WaitGroup
	insert := func() {
		for i := 0; i < 10; i++ {
			q.Insert(i)
		}
		wg.Done()
	}
	remove := func() {
		for i := 0; i < 10; i++ {
			_ = q.Remove()
		}
		wg.Done()
	}
	const n = 10
	wg.Add(2 * n)
	// start 2n goroutines:
	// - n insert goroutines
	// - n remove goroutines
	for i := 0; i < n; i++ {
		go insert()
		go remove()
	}
	// wait for 2n goroutines to finish
	wg.Wait()
}

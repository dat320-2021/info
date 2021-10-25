package concurrentlist

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentList(t *testing.T) {
	l := New()
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numOps; j++ {
				l.Insert(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numOps; j++ {
				n := l.LookupNodeLock(j)
				if n.key != j {
					t.Errorf("l.Lookup(%d) = %d, expected %d", j, n.key, j)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

var l = New()

const (
	numGoroutines = 10
	numOps        = 100
)

func BenchmarkConcurrentListLookup(b *testing.B) {
	var n *Node
	for ops := 1; ops < 1024; ops = 2 * ops {
		b.Run(fmt.Sprintf("len=%d", ops), func(b *testing.B) {
			for j := 0; j < ops; j++ {
				l.Insert(j)
			}
			for k := 0; k < b.N; k++ {
				var wg sync.WaitGroup
				wg.Add(numGoroutines)
				for i := 0; i < numGoroutines; i++ {
					go func() {
						for j := 0; j < numOps; j++ {
							n = l.LookupNodeLock(ops)
						}
						wg.Done()
					}()
				}
				wg.Wait()
			}
			result = n
		})
	}
}

var result *Node

func BenchmarkConcurrentListInsert(b *testing.B) {
	for k := 0; k < b.N; k++ {
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				for j := 0; j < numOps; j++ {
					l.Insert(j)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

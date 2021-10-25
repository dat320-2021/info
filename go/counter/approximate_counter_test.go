package counter

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"

	"golang.org/x/tools/benchmark/parse"
)

func grabBenchmark(s string) {
	for _, line := range strings.Split(s, "\n") {
		bench, err := parse.ParseLine(line)
		if err != nil {
			// ignore lines that are not benchmark lines
			continue
		}
		_ = bench
	}
}

func BenchmarkConcurrentApproxCounter(b *testing.B) {
	for threshold := 1; threshold <= 256; threshold *= 2 {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runLoop(threshold)
			}
		})
	}
}

func BenchmarkConcurrentApproxCounterNoLock(b *testing.B) {
	for threshold := 1; threshold <= 256; threshold *= 2 {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runLoopNoLock(threshold)
			}
		})
	}
}

const totCount = 1_000_000

func runLoop(threshold int) {
	numCPUs := runtime.NumCPU()
	c := New(threshold, numCPUs)
	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < totCount; j++ {
				c.Update(id, 1)
			}
		}(i)
	}
	wg.Wait()
}

func runLoopNoLock(threshold int) {
	numCPUs := runtime.NumCPU()
	c := New(threshold, numCPUs)
	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < totCount; j++ {
				c.UpdateNoLocalLock(id, 1)
			}
		}(i)
	}
	wg.Wait()
}

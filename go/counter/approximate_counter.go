package counter

import "sync"

type ApproximateCounter struct {
	global    int
	glock     sync.Mutex
	local     []int
	llock     []sync.Mutex
	threshold int
}

func New(threshold, numLocalCounters int) *ApproximateCounter {
	return &ApproximateCounter{
		threshold: threshold,
		local:     make([]int, numLocalCounters),
		llock:     make([]sync.Mutex, numLocalCounters),
	}
}

func (c *ApproximateCounter) Update(id, amount int) {
	c.llock[id].Lock()
	defer c.llock[id].Unlock()
	c.local[id] += amount
	if c.local[id] >= c.threshold {
		// transfer local count value to global counter
		c.glock.Lock()
		c.global += c.local[id]
		c.glock.Unlock()
		c.local[id] = 0
	}
}

func (c *ApproximateCounter) Get() int {
	c.glock.Lock()
	defer c.glock.Unlock()
	return c.global
}

// UpdateNoLocalLock updates the approximate counter without locking the local counters.
// This is safe as long as only one goroutine updates the local counter.
func (c *ApproximateCounter) UpdateNoLocalLock(id, amount int) {
	c.local[id] += amount
	if c.local[id] >= c.threshold {
		// transfer local count value to global counter
		c.glock.Lock()
		c.global += c.local[id]
		c.glock.Unlock()
		c.local[id] = 0
	}
}

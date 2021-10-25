package counter

import "sync"

// Concurrent counters

// Sometimes called reference counters.

type counter struct {
	mu  sync.Mutex
	cnt int
}

func NewCounter() *counter {
	return &counter{}
}

func (c *counter) increment() {
	c.mu.Lock()
	c.cnt++ // critical section
	c.mu.Unlock()
}

func (c *counter) noLockInc() {
	c.cnt++ // critical section
}

func (c *counter) decrement() {
	c.mu.Lock()
	c.cnt-- // critical section
	c.mu.Unlock()
}

func (c *counter) setVal(v int) {
	c.mu.Lock()
	c.cnt = v // critical section
	c.mu.Unlock()
}

func (c *counter) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cnt // critical section
}

func (c *counter) get2() int {
	c.mu.Lock()
	tmp := c.cnt // critical section
	c.mu.Unlock()
	return tmp
}

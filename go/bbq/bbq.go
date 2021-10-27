package bbq

import "sync"

const MAX = 10

// BBQueue is a bounded buffer queue.
type BBQueue struct {
	// Synchronization variables
	lock        sync.Mutex
	itemAdded   sync.Cond
	itemRemoved sync.Cond
	// State variables
	items     [MAX]int
	front     int
	nextEmpty int
}

func New() *BBQueue {
	q := &BBQueue{}
	q.itemAdded = *sync.NewCond(&q.lock)
	q.itemRemoved = *sync.NewCond(&q.lock)
	return q
}

// Insert item in queue. Insert blocks if queue is full, waiting for
// a corresponding Remove operation to make room for this item.
func (q *BBQueue) Insert(item int) {
	q.lock.Lock()
	// begin critical section
	for q.nextEmpty-q.front == MAX {
		q.itemRemoved.Wait() // queue is full, wait for item to be removed
	}
	q.items[q.nextEmpty%MAX] = item
	q.nextEmpty++
	q.itemAdded.Signal()
	// end critical section
	q.lock.Unlock()
}

// Remove an item from queue. Remove blocks if the queue is empty,
// waiting for corresponding Insert operations to add items that
// can subsequently be removed.
func (q *BBQueue) Remove() (item int) {
	q.lock.Lock()
	// begin critical section
	for q.front == q.nextEmpty {
		q.itemAdded.Wait() // queue is empty, wait for an item to be added
	}
	item = q.items[q.front%MAX]
	q.front++
	q.itemRemoved.Signal()
	// end critical section
	q.lock.Unlock()
	return
}

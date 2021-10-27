package bbq

// ChannelQueue has identical behavior to BBQueue.
type ChannelQueue struct {
	items chan int
}

func NewChannelQueue() *ChannelQueue {
	return &ChannelQueue{
		items: make(chan int, MAX),
	}
}

// Insert item in queue. Insert blocks if queue is full, waiting for
// a corresponding Remove operation to make room for this item.
func (q *ChannelQueue) Insert(item int) {
	q.items <- item
}

// Remove an item from queue. Remove blocks if the queue is empty,
// waiting for corresponding Insert operations to add items that
// can subsequently be removed.
func (q *ChannelQueue) Remove() int {
	return <-q.items
}

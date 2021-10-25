package concurrentlist

import "sync"

type Node struct {
	sync.Mutex
	next *Node
	key  int
}

type List struct {
	sync.Mutex
	head *Node
}

func New() *List {
	return &List{}
}

func (l *List) Insert(key int) {
	l.Lock()
	defer l.Unlock()
	newNode := &Node{key: key, next: l.head}
	l.head = newNode
}

func (l *List) Lookup(key int) *Node {
	l.Lock()
	defer l.Unlock()
	for n := l.head; n != nil; n = n.next {
		if n.key == key {
			return n
		}
	}
	return nil
}

func (l *List) LookupNodeLock(key int) *Node {
	l.Lock()
	n := l.head // fast critical section
	l.Unlock()
	for n != nil {
		n.Lock()
		if n.key == key {
			n.Unlock()
			return n
		}
		tmp := n.next
		n.Unlock()
		n = tmp
	}
	return nil
}

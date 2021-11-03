package main

import (
	"sync"
)

type Node struct {
	next *Node
	key  int
}

type List struct {
	sync.Mutex
	head *Node
}

func (l *List) Insert(key int) {
	l.Lock()
	defer l.Unlock()
	newNode := &Node{key: key, next: l.head}
	l.head = newNode
}

// NOTE: This is not a proper atomic insert.
// It is only meant to illustrate the goings-on.
// It will not work as an atomic insert.

func (l *List) InsertCAS(key int) {
	newNode := &Node{key: key}
	newNode.next = l.head
	for !CAS2(l.head, *newNode.next, *newNode) {
		newNode.next = l.head
	}
}

func CAS2(address *Node, expected, new Node) bool {
	if *address == expected {
		*address = new
		return true
	}
	return false
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

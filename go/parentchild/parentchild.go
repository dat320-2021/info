package main

import (
	"fmt"
	"info/go/semaphore"
)

const X = 0

var sem = semaphore.New(X)

func child(arg string) {
	fmt.Println(arg)
	// signal: child is done
	sem.Post()
}

func main() {
	fmt.Println("parent: begin")
	go child("A")
	sem.Wait() // wait for child to finish
	fmt.Println("parent: end")
}

package main

import "fmt"

func x() {
	var k int // TODO: accesses should be protected with lock
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			k = i * 10
		}()
	}
	fmt.Println(k)
}

func x2() {
	var k int // TODO: accesses should be protected with lock
	for i := 0; i < 10; i++ {
		j := i
		go func() {
			k = j * 10
		}()
	}
	fmt.Println(k)
}

func y() {
	var k int // TODO: accesses should be protected with lock
	for i := 0; i < 10; i++ {
		go func(i int) {
			k = i * 10
		}(i)
	}
	fmt.Println(k)
}

func y2() {
	var k int // TODO: accesses should be protected with lock
	for i := 0; i < 10; i++ {
		go func(j int) {
			k = j * 10
		}(i)
	}
	fmt.Println(k)
}

// Two examples to illustrate that we need to make a copy of
// the loop variable i when we want to use it within a goroutine.
// There are two ways to do this. One way, as in x(), is to copy
// the loop variable i to a new local i variable (Line 8); that is,
// Line 8 creates a new variable for each loop iteration, copied
// from the loop variable i. The approach in x2() is the same, except
// here the local variable is renamed to j to make it clearer what's
// happening.
//
// The other approach, as shown in y(), is to pass the loop variable i
// as an argument to the goroutine anonymous function. Similar to the
// x2() function, we can also rename it to j as in y2() to make it
// clearer what's happening..
//
// The use of the func() literal above are example uses of
// anonymous functions, also called a closure.
//
// See also: https://golang.org/doc/faq#closures_and_goroutines

func main() {
	x()
	y()
}

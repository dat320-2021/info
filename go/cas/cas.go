package main

import "fmt"

// NOTE: This is not a proper AtomicInc.
// It is only meant to illustrate the goings-on.
// It will not work as an atomic increment.

func AtomicInc(value *int, amount int) {
	old := *value
	for !CAS(value, old, old+amount) {
		old = *value
	}
}

func CAS(address *int, expected, new int) bool {
	if *address == expected {
		*address = new
		return true
	}
	return false
}

func main() {
	val := 9
	fmt.Println("val=", val)
	AtomicInc(&val, 5)
	fmt.Println("val=", val)
}

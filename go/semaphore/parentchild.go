package semaphore

import "fmt"

const X = 0

var sem = &Semaphore{Value: X}

func child(arg string) {
	fmt.Println(arg)
	// signal: child is done
	sem.Post()
}

func mainX() {
	fmt.Println("parent: begin")
	go child("A")
	sem.Wait() // wait for child to finish
	fmt.Println("parent: end")
}

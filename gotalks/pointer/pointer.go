package main

import "fmt"

func main() {
	a := "hello"
	p := &a
	fmt.Printf("%v\n", p)
	fmt.Println(*p)
	fmt.Println(string(*p) + ", dat320 世界")
}

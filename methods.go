package main

import "fmt"

type rect struct {
	width, length int
}

// pointer not required, ust for example
// here area() has a reciever type of *rect
func (r *rect) area() int {
	return r.length * r.width
}

func (r rect) perim() int {
	return 2 * (r.length + r.width)
}

func main() {
	r1 := rect{1, 2}
	fmt.Println(r1)
	fmt.Println(r1.area())
	// doesn't work fmt.Println(area(r1))
}

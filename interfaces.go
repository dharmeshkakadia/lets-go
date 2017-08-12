package main

import "fmt"
import "math"

type geommetry interface {
	area() float64
}

type rect struct {
	length float64
	width  float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getarea(g geommetry) float64 {
	return g.area()
}

func main() {
	r1 := rect{length: 4, width: 5}
	fmt.Println(getarea(r1))
	fmt.Println(getarea(circle{5}))
}

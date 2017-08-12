package main

import "fmt"

func zero(a int) {
	a = 0
}

func zeropt(a *int) {
	*a = 0
}

func main() {
	a := 1
	zero(a)
	fmt.Println(a)
	zeropt(&a)
	fmt.Println(a)
}

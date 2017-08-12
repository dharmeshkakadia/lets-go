package main

import "fmt"

func myfunc(a int, b int) (int, int, int) {
	return a, b, a + b
}

func main() {
	a, _, c := myfunc(1, 2)
	fmt.Println(a, c)
}

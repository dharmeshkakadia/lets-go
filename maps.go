package main

import "fmt"

func main() {
	a := map[string]int{"k2": 7}
	b := make(map[string]int)

	a["k1"] = 1
	fmt.Println(a, b)

	delete(a, "k1")
	delete(b, "k1")
	fmt.Println(a, b)

	_, isPresent := a["k1"]
	fmt.Println(isPresent)
}

package main

import "fmt"

func main() {
	a := make([]string, 3)
	fmt.Println(a)

	a[1] = "s"
	fmt.Println(a)

	a = append(a, "blah")
	fmt.Println(a)
	cp := make([]string, len(a))
	copy(cp, a)
	fmt.Println(cp)

	subslice := a[1:2]
	fmt.Println(subslice, a[1:], a[:4])

	directslice := []string{"a", "b"}
	fmt.Println(directslice)
}

package main

import "fmt"

func f(a int, b ...string) int {
	return len(b)
}

func main() {
	fmt.Println(f(1, "a"))
	fmt.Println(f(1, "a", "b", "c"))

	// slices integration
	s := []string{"z", "y", "x", "j"}
	fmt.Println(f(1000, s...))
}

package main

import "fmt"

func intSeq(add int) func() int {
	i := 0
	return func() int {
		i += add
		return i
	}
}

func main() {
	nextInt := intSeq(1)
	fmt.Println(nextInt())
	fmt.Println(nextInt()) // we get the next number since we have pointer to the inter function
	fmt.Println(nextInt())

	// restart the sequence
	nextInt = intSeq(1)
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	evenInts := intSeq(2)
	fmt.Println(evenInts())
	fmt.Println(evenInts())

}

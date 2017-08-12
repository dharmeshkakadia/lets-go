package main

import "fmt"
import "sort"

func main() {
	s := []string{"c", "b", "a"}
	sort.Strings(s)
	fmt.Println(s)

	n := []int{1, 3, 2}
	sort.Ints(n)
	fmt.Println(n)
	fmt.Println(sort.IntsAreSorted(n))
}

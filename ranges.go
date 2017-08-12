package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	sum := 0
	sumIndex := 0
	for i, num := range a {
		sum += num
		sumIndex += i
	}
	fmt.Println(sum, sumIndex)

	m := map[int]string{1: "v1", 2: "v2", 3: "v3"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// only iterate over keys
	for k := range m {
		fmt.Println(k)
	}

	// string range is unicode values
	for _, s := range "blah" {
		fmt.Println(s)
	}
}

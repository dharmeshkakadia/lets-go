package main

import "fmt"

func main() {
	var a [6]int
	fmt.Println(a)

	a[5] = 5
	fmt.Println(a[5])
	fmt.Println(len(a))

	b := [3]int{4, 7, 8}
	fmt.Println(b)

	// does not work : twoC := [][]int{1, 2, 3, 4, 5, 6}
	twoD := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(twoD)
	for i := 0; i < len(twoD); i++ {
		fmt.Println(twoD[i])
	}
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			fmt.Println(twoD[i][j])
		}
	}
}

package main

import "fmt"

func main() {
	n := 2
	if n < 4 {
		fmt.Println("bingo")
	} else {
		fmt.Println("whatever")
	}

	// if blocks can have variables
	if v := 5; v < 3 {
		fmt.Println("wrong")
	} else if v < 6 {
		fmt.Println(v)
	}

}

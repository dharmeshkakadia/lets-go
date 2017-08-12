package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("boom")
		break
	}
}

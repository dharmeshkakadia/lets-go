package main

import "fmt"

func main() {
	c := make(chan string, 5)
	c <- "one"
	c <- "two"
	close(c)

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println()
}

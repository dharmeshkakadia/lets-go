package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "first"
	messages <- "second"
	// wont work since this would deadlock : messages <- "third"
	fmt.Println(<-messages, "and", <-messages)
}

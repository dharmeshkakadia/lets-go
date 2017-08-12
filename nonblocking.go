package main

import "fmt"

func main() {
	msg := make(chan string)
	singals := make(chan bool)
	select {
	case m := <-msg:
		fmt.Println(m)
	default:
		fmt.Println("no messages received")
	}

	select {
	case msg <- "hi":
		fmt.Println("sent sucessful")
	default:
		fmt.Println("no message sent")
	}

	select {
	case m := <-msg:
		fmt.Println("received message ", m)
	case s := <-singals:
		fmt.Println("received singal ", s)
	default:
		fmt.Println("Nothing here")
	}

}

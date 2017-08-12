package main

import "fmt"

func ping(sendingChannel chan<- string, message string) {
	sendingChannel <- message
}

func pong(recievingChannel <-chan string, newSendingChannel chan<- string) {
	msg := <-recievingChannel
	newSendingChannel <- msg + "pong"
}

func pong2(recievingChannel <-chan string) chan string {
	newChannel := make(chan string, 1)
	newChannel <- <-recievingChannel + "pong2" // msg := <-recievingChannel ; newChannel <- msg + "pong2"
	return newChannel
}

func main() {
	pipe1 := make(chan string, 1)
	pipe2 := make(chan string, 1)
	ping(pipe1, "ping")
	pong(pipe1, pipe2)
	fmt.Println(<-pipe2)
	ping(pipe1, "ping")

	fmt.Println(<-pong2(pipe1))
}

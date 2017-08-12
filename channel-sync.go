package main

import "fmt"
import "time"

func slowAsyncFunc(pipe chan bool) {
	fmt.Println("Doing some busy work")
	time.Sleep(time.Second)
	fmt.Println("done")
	pipe <- true
}

func main() {
	pipe := make(chan bool, 1)
	fmt.Println("starting the work")
	go slowAsyncFunc(pipe)
	fmt.Println("continue doing something else")
	<-pipe
}

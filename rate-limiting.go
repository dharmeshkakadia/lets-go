package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Second) //time.Tick is convience wrapper that directly gives the ticker channel
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) // burst of 3 requests are allowed

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() // fill the channel with burstable capacity. Note that this only intial burst support and once exhausted it will go back to being ticker
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 10)
	for i := 0; i < 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for r := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", r, time.Now())
	}
	fmt.Println()
}

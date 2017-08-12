package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	wait := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				// process job
				fmt.Println("got job ", j)
			} else {
				fmt.Println("Done with all the jobs")
				wait <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job ", i)
	}
	close(wait)
	fmt.Println("sent all jobs")
	<-wait
}

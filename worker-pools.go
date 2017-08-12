package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println(id, "received job :", j)
		time.Sleep(time.Second)
		fmt.Println(id, "sending result for job :", j)
		results <- j
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	jobs <- 1
	jobs <- 2
	jobs <- 3

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	jobs <- 4
	jobs <- 5
	close(jobs)

	// following will exit before getting all the results
	// for r := range results {
	// 	fmt.Println("got results", r)
	// }

	for r := 0; r < 5; r++ {
		fmt.Println("got result", <-results)
	}
}

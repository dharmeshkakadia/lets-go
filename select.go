package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "m1"

	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "m2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("from c1 : ", msg1)
		case msg2 := <-c2:
			fmt.Println("from c2 : ", msg2)
		}
	}

	fmt.Println()
}

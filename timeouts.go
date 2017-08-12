package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "m1"
	}()

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "m2"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout")
	}

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout 2!")
	}

	fmt.Println()
}

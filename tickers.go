package main

import "fmt"
import "time"

func main() {
	ticker1 := time.NewTicker(time.Second)

	go func() {
		for i := range ticker1.C {
			fmt.Println("got a tick : ", i)
		}
	}()

	time.Sleep(time.Second * 4)
	ticker1.Stop()
	fmt.Println()
}

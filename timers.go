package main

import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer one ran down")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		fmt.Println("not timing yet")
		<-timer2.C
		fmt.Println("Timer2 Expired !") // will not happen because it will be stopped before it expires
	}()
	time.Sleep(time.Second)
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("Timer2 is stopped")
	}
	fmt.Println()
}

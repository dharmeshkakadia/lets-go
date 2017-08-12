package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("waiting for signal (use ctrl-c)")
	<-done
	fmt.Println("exiting")
}

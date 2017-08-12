package main

import "fmt"

func f(num int, message string) {
	for i := 0; i < num; i++ {
		fmt.Println(message, ":", i)
	}
}

func main() {
	f(5, "sync")
	go f(3, "async")

	go func(a int) {
		for i := 0; i < a; i++ {
			fmt.Println("inline async : ", i)
		}
	}(2)
	// have to wait for async tasks. Without this program will exit before the sync finishes.
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Are you saying : %s ?\n", input)
}

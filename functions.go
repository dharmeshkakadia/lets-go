package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// can specify type only once if all are same
func sum3(a, b, c int, message string) int {
	fmt.Println(message)
	return a + b + c
}

func main() {
	fmt.Println("sum : ", sum(1, 2))
	fmt.Println("sum : ", sum3(1, 2, 4, "sum3 is called"))
}

package main

import "os"
import "fmt"

func main() {
	defer fmt.Println("Hello") // this will not aclled
	os.Exit(3)
}

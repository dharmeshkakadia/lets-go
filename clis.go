package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)     //including programname
	fmt.Println(os.Args[1:]) // only arguments
	fmt.Println(os.Args[2])  // second argument
}

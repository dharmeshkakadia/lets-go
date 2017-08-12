package main

import (
	"fmt"
	"os"
)

func main() {
	p := os.Setenv("foo", "bar") // return value is previous value
	fmt.Println(p, "is reset to", os.Getenv("foo"))

	for _, e := range os.Environ() { //ignore the index
		fmt.Println(e)
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	wordptr := flag.String("word", "defaultWord", "Plese provide a word") // -word "value"
	flag.Parse()
	fmt.Println(*wordptr)
	fmt.Println(flag.Args()) // remaining non-flag arguments. They have to be at the end
}

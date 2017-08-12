package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() { // one line at a time
		us := strings.ToUpper(s.Text())
		fmt.Println(us)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "erorr:", err)
		os.Exit(1)
	}
}

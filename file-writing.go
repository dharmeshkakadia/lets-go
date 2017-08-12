package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	e := ioutil.WriteFile("/tmp/file2.txt", []byte("hello\n"), 0644)
	check(e)
	f, err := os.Create("/tmp/file1.txt")
	check(err)
	defer f.Close()
	f.Write([]byte("content\n"))
	f.WriteString("other string\n")
	f.Sync()
	w := bufio.NewWriter(f)
	w.WriteString("more string\n")
	w.Flush()
	fmt.Println()
}

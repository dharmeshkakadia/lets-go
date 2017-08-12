package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
	f, err := os.Open("/tmp/dat")
	check(err)
	b := make([]byte, 5)
	r, err := f.Read(b)
	check(err)
	fmt.Println("bytes read", r)
	fmt.Println("contents", string(b))
	f.Seek(6, 0)
	// b = make([]byte, 5)
	f.Read(b)
	fmt.Println("content after seek", string(b))

	n, err := io.ReadAtLeast(f, b, 5)
	check(err)
	fmt.Println(n)
	fmt.Println(string(b))

	r2 := bufio.NewReader(f)
	b, err = r2.Peek(5)
	check(err)
	fmt.Println(string(b))
	f.Close()
}

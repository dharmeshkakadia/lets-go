package main

import "os"
import "fmt"

//defer == finalize in java

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f, "data")
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("writing")
	fmt.Fprintln(f, data)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

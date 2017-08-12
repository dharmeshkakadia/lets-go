package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	if o, err := dateCmd.Output(); err == nil {
		fmt.Println(string(o))
	}

	grepCmd := exec.Command("grep", "hello")
	in, _ := grepCmd.StdinPipe()
	out, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	in.Write([]byte("sdfsd hello ksjdhs \n dsfkjdslfs grlsk"))
	in.Close()

	o, _ := ioutil.ReadAll(out)
	fmt.Println(string(o))
	grepCmd.Wait()

	if o, err := exec.Command("bash", "-c", "ls -l -a -h").Output(); err == nil {
		fmt.Println(string(o))
	}
}

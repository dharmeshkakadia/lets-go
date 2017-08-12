package main

import "fmt"

type s struct {
	name    string
	address string
}

func main() {
	fmt.Println(s{"f1", "l1"})
	fmt.Println(s{address: "l2", name: "f1"}) // fmt.Println(s{address: "l2", "f1"}) doesnt work
	fmt.Println(s{address: "l2"})
	s := s{name: "name1", address: "address1"} // type s and variable s is diff :)
	fmt.Println(s)
	fmt.Println(s.name)
	sp := &s
	sp.address = "WA"
	fmt.Println(s.address)
}

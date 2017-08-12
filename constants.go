package main

import "fmt"
import "math"

const con string = "Blah"
const con2 = "foo"

var con3 = "awe"

// con4:="boom"

func main() {
	fmt.Println("con,con2,con3", con, con2, con3)
	const con = "ggg"
	fmt.Println(con)

	//Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / 50000000
	fmt.Println(d)

	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int(d))
	fmt.Println(int64(d))
	// auto conversion to float
	fmt.Println(math.Sin(d))
}

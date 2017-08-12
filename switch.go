package main

import "fmt"
import "time"

func main() {
	v := "value"

	switch v {
	case "value":
		fmt.Println(v)
	default:
		fmt.Println("default")
	}

	// runtime values, multiple conditions in switch,case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println()
	}

	//switch without condition is like if-else
	t := time.Now()
	switch {
	//dynamic condition in case
	case t.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an int type")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a bool")
		default:
			fmt.Printf("I dont know %T: ", t)
		}
	}

	whatAmI("hellow")
	whatAmI(true)
	whatAmI(5)
	whatAmI(6 / 7)
	whatAmI(6.0 / 7)
}

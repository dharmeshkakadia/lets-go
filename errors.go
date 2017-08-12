package main

import "fmt"
import "errors"

func funcWithError(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("can't divide by zero")
	}
	return a / b, nil
}

type detailedError struct {
	code    int
	message string
}

func (e *detailedError) Error() string {
	return fmt.Sprintf("Error code:%d, failed with : %s", e.code, e.message)
}

func funcWithDeatiledError(a, b float64) (float64, error) {
	if b == 0 {
		return -1, &detailedError{code: -1, message: "Can't divide by zero"}
	}
	return a / b, nil
}
func main() {
	fmt.Println(funcWithError(1, 2))
	fmt.Println(funcWithError(1, 0))
	for _, i := range []float64{-1, 0, 0.5, 1, 2} {
		if returnValue, error := funcWithError(1, i); error == nil {
			fmt.Println("Sucess ", returnValue, error)
		} else {
			fmt.Println("Failed ", returnValue, error)
		}
	}
	for _, i := range []float64{-1, 0, 0.5, 1, 2} {
		if returnValue, error := funcWithDeatiledError(1, i); error == nil {
			fmt.Println("Sucess ", returnValue, error)
		} else {
			fmt.Println("Failed ", returnValue, error.(*detailedError).code)
		}
	}
}

package main

import (
	"errors"
	"fmt"
	"time"
)

type Errors struct {
	msg      string
	function string
	time     time.Time
}

func (e *Errors) Error() string {
	return e.msg
}

func (e *Errors) GetHost() time.Time {
	return e.time
}

func main() {
	errs := []interface{}{
		Errors{"Custom error.", "Error function", time.Now()},
		errors.New("Generic error."),
	}

	for _, e := range errs {
		if v, ok := e.(Errors); ok {
			fmt.Println("Error: ", v.Error())
			fmt.Println("Type: Custom Error.")
		} else {
			fmt.Println("Error: ", v.Error())
			fmt.Println("Type: General Error.")
		}
	}
}

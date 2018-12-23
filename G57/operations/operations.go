package operations

import (
	"errors"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("division by zero is undefined")
	}
	return a / b, nil
}

func Quit() {
	os.Exit(1)
}

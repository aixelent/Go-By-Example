package calc

import (
	"Go-By-Example/G57/operations"
	"fmt"
)

func Calculate(a, b int, operation string) (string, int) {
	switch operation {
	case "+":
		return "Add operation:", operations.Add(a, b)
	case "-":
		return "Subtract operation:", operations.Subtract(a, b)
	case "*":
		return "Multiply operation:", operations.Multiply(a, b)
	case "/":
		result, _ := operations.Divide(a, b)
		return "Divide operation:", result
	case ":q":
		operations.Quit()
	}
	fmt.Println("Invalid operation.")
	return "Choose operator from: '+', '-', '*', /'", 0
}

package main

import (
	"fmt"
	"goprojects/calc"
)

func main() {
	a := 20.0
	b := 10.0

	fmt.Println("Addition:", calc.Add(a, b))
	fmt.Println("Subtraction:", calc.Subtract(a, b))
	fmt.Println("Multiplication:", calc.Multiply(a, b))

	result, err := calc.Divide(a, b)
	if err != nil {
		fmt.Println("Error in division:", err)
	} else {
		fmt.Println("Division:", result)
	}
}

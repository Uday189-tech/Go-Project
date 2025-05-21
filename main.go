package main

import (
	"fmt"
	"goprojects/calc"
)

func main() {
	a := 5
	b := 10

	sum := calc.Add(a, b)

	fmt.Println("Sum:", sum)
}

package main

import (
	"fmt"

	"labassignment/main/custompkg"
)

func main() {

	fmt.Println("===== CUSTOM PACKAGE DEMONSTRATION =====")

	// String functions
	text := "Elsa"

	fmt.Println("Original String:", text)
	fmt.Println("Reversed String:", custompkg.Reverse(text))
	fmt.Println("Uppercase String:", custompkg.ToUpper(text))

	// Mathematical functions
	fmt.Println("\n===== MATHEMATICAL FUNCTIONS =====")

	fmt.Println("Factorial of 5:", custompkg.Factorial(5))
	fmt.Println("Power of 2^3:", custompkg.Power(2, 3))
}

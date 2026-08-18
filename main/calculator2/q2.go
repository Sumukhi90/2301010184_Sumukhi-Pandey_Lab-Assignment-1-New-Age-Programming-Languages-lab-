package main

import "fmt"

func main() {
	var a, b int
	var x, y float64

	for {
		fmt.Print("Enter two integers: ")
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			fmt.Println("Invalid input, please enter numbers only.")
			continue
		}
		if a < 0 || b < 0 {
			fmt.Println("Values out of range, please enter non-negative integers.")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter two floats: ")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			fmt.Println("Invalid input, please enter numbers only.")
			continue
		}
		if x < 0 || y < 0 {
			fmt.Println("Values out of range, please enter non-negative numbers.")
			continue
		}
		break
	}

	fmt.Println("Integer Addition:", a+b)
	fmt.Println("Integer Subtraction:", a-b)
	fmt.Println("Integer Multiplication:", a*b)

	fmt.Println("Float Addition:", x+y)
	fmt.Println("Float Subtraction:", x-y)
	fmt.Println("Float Multiplication:", x*y)
}

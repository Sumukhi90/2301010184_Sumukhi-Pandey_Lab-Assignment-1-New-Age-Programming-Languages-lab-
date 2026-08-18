package main

import "fmt"

func main() {
	var a, b int = 10, 3
	var x, y float64 = 10.5, 3.2

	fmt.Println("Integer Addition:", a+b)
	fmt.Println("Integer Subtraction:", a-b)
	fmt.Println("Integer Multiplication:", a*b)

	fmt.Println("Float Addition:", x+y)
	fmt.Println("Float Subtraction:", x-y)
	fmt.Println("Float Multiplication:", x*y)
}

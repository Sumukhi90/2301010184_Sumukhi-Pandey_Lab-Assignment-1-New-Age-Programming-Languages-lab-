package main

import "fmt"

// Function that modifies the original variable using a pointer
func modifyValue(n *int) {
	*n = *n + 10
}

// Structure for demonstrating new()
type Student struct {
	Name string
	Age  int
}

func main() {

	// 1. Referencing and Dereferencing
	x := 20

	fmt.Println("----- Pointer Referencing and Dereferencing -----")
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value using pointer:", *(&x))

	// 2. Passing pointer to a function
	fmt.Println("\n----- Pass-by-Reference Using Pointer -----")
	fmt.Println("Before modification:", x)

	modifyValue(&x)

	fmt.Println("After modification:", x)

	// 3. Allocating a struct using new()
	fmt.Println("\n----- Struct Using new() -----")

	student := new(Student)

	// Accessing and modifying fields through pointer
	student.Name = "Sumukhi"
	student.Age = 21

	fmt.Println("Student Name:", student.Name)
	fmt.Println("Student Age:", student.Age)

	// Modifying struct field through pointer
	student.Age = 22

	fmt.Println("Modified Student Age:", student.Age)
}

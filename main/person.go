package main

import "fmt"

// Person structure
type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

// Method to read data into the Person structure
func (p *Person) readData() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

// Method to display all Person details
func (p Person) displayData() {
	fmt.Println("\n----- Person Details -----")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Println("Salary:", p.Salary)
}

func main() {

	// Creating first Person object
	var person1 Person

	fmt.Println("Enter details for Person 1:")
	person1.readData()
	person1.displayData()

	// Creating second Person object
	var person2 Person

	fmt.Println("\nEnter details for Person 2:")
	person2.readData()
	person2.displayData()
}

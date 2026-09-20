package main

import "fmt"

func main() {

	// =====================================
	// PART 1: SLICE OPERATIONS
	// =====================================

	fmt.Println("===== SLICE OPERATIONS =====")

	// Initial slice
	students := []string{"Sumukhi", "Ananya", "Riya"}

	fmt.Println("Initial Slice:", students)

	// ADD OPERATION
	students = append(students, "Priya")

	fmt.Println("After Adding Priya:", students)

	// REMOVE OPERATION (Remove index 1)
	index := 1

	students = append(students[:index], students[index+1:]...)

	fmt.Println("After Removing Index 1:", students)

	// UPDATE OPERATION
	students[1] = "Neha"

	fmt.Println("After Updating Index 1:", students)

	// =====================================
	// PART 2: MAP OPERATIONS
	// =====================================

	fmt.Println("\n===== MAP OPERATIONS =====")

	// Initial map
	marks := map[string]int{
		"Maths":   85,
		"Science": 90,
	}

	fmt.Println("Initial Map:", marks)

	// INSERT OPERATION
	marks["English"] = 88

	fmt.Println("After Inserting English:", marks)

	// DELETE OPERATION
	delete(marks, "Science")

	fmt.Println("After Deleting Science:", marks)

	// LOOKUP OPERATION
	subject := "Maths"

	mark, exists := marks[subject]

	if exists {
		fmt.Println("Lookup Maths:", mark)
	} else {
		fmt.Println("Subject not found")
	}
}

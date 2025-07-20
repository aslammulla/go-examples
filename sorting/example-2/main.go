package main

import (
	"fmt"
	"sort"
)

// Define the Employee struct with Department
type Employee struct {
	Name       string
	Department string
	Salary     int
}

func main() {
	employees := []Employee{
		{Name: "Aslam", Department: "Engineering", Salary: 70000},
		{Name: "Ravi", Department: "Engineering", Salary: 90000},
		{Name: "Jitesh", Department: "Engineering", Salary: 50000},
		{Name: "Komal", Department: "Testing", Salary: 80000},
		{Name: "Harsha", Department: "Product", Salary: 75000},
	}

	fmt.Println("Before sorting (by department → salary):")
	for _, e := range employees {
		fmt.Printf("%s (%s): %d\n", e.Name, e.Department, e.Salary)
	}

	// Sort by Department ASC, then Salary DESC
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Department == employees[j].Department {
			return employees[i].Salary > employees[j].Salary
		}
		return employees[i].Department < employees[j].Department
	})

	fmt.Println("\nAfter sorting (by department asc → salary desc):")
	for _, e := range employees {
		fmt.Printf("%s (%s): %d\n", e.Name, e.Department, e.Salary)
	}
}

/*
OUTPUT:
$ go run main.go
Before sorting (by department → salary):
Aslam (Engineering): 70000
Ravi (Engineering): 90000
Jitesh (Engineering): 50000
Komal (Testing): 80000
Harsha (Product): 75000

After sorting (by department asc → salary desc):
Ravi (Engineering): 90000
Aslam (Engineering): 70000
Jitesh (Engineering): 50000
Harsha (Product): 75000
Komal (Testing): 80000
*/

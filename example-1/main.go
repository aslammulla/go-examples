package main

import (
	"fmt"
	"sort"
)

// Define the Employee struct
type Employee struct {
	Name   string
	Salary int
}

// Create a new type to implement sort.Interface
type BySalary []Employee

func (e BySalary) Len() int           { return len(e) }
func (e BySalary) Less(i, j int) bool { return e[i].Salary > e[j].Salary } // Descending
func (e BySalary) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func main() {
	employees := []Employee{
		{Name: "Aslam", Salary: 70000},
		{Name: "Ravi", Salary: 90000},
		{Name: "Jitesh", Salary: 50000},
		{Name: "Komal", Salary: 80000},
	}

	fmt.Println("Before sorting (by salary):")
	for _, e := range employees {
		fmt.Printf("%s: %d\n", e.Name, e.Salary)
	}

	sort.Sort(BySalary(employees))

	fmt.Println("\nAfter sorting (by salary descending):")
	for _, e := range employees {
		fmt.Printf("%s: %d\n", e.Name, e.Salary)
	}
}

/*
OUTPUT:
$ go run main.go
Before sorting (by salary):
Aslam: 70000
Ravi: 90000
Jitesh: 50000
Komal: 80000

After sorting (by salary descending):
Ravi: 90000
Komal: 80000
Aslam: 70000
Jitesh: 50000
*/

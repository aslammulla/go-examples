package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Create a new validator instance
	validate := validator.New()
	// Example of validating a single field: an email address
	email := "invalid-email@"
	err := validate.Var(email, "required,email")
	if err != nil {
		fmt.Println("Invalid email address!")
		return
	}

	fmt.Println("Email is valid!")
}

/*
OUTPUT:
$ go run main.go
Invalid email address!
*/

package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// User Registration Struct with Built-in Validation Tags
type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"` // Must be 3-20 alphanumeric characters
	Email    string `json:"email" validate:"required,email"`                    // Must be a valid email
	Age      int    `json:"age" validate:"required,gte=18,lte=60"`              // Must be between 18 and 60
	Phone    string `json:"phone" validate:"required,len=10,numeric"`           // Must be a 10-digit number
	Website  string `json:"website" validate:"omitempty,url"`                   // Optional but must be a valid URL if provided
}

func main() {
	validate := validator.New()

	// Simulated Invalid Input
	user := RegisterUserRequest{
		Username: "as",           // Too short
		Email:    "invalidemail", // Not a valid email
		Age:      17,             // Less than 18
		Phone:    "123abc",       // Not a 10-digit number
		Website:  "not-a-url",    // Invalid URL
	}

	// Run Validation
	if err := validate.Struct(user); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on '%s' validation\n", e.Field(), e.Tag())
		}
		return
	}

	fmt.Println("All fields are valid!")
}

/*
OUTPUT:
$ go run main.go
Field 'Username' failed on 'min' validation
Field 'Email' failed on 'email' validation
Field 'Age' failed on 'gte' validation
Field 'Phone' failed on 'len' validation
Field 'Website' failed on 'url' validation
*/

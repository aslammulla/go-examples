package main

import (
	"fmt"

	"github.com/aslammulla/go-examples/validation/example3/validation"
	"github.com/go-playground/validator/v10"
)

// Struct with built-in + custom tags
type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,notreserved"` // custom
	Password string `json:"password" validate:"required,strongpwd"`   // custom
	Phone    string `json:"phone"    validate:"required,phone"`       // custom
	Email    string `json:"email"    validate:"required,email"`       // built-in
	Age      int    `json:"age"      validate:"gte=18,lte=60"`        // built-in
}

func main() {
	validate := validator.New()

	// Register custom validation rules
	validation.RegisterCustomRules(validate)

	// Simulate invalid input
	user := RegisterUserRequest{
		Username: "admin",    // reserved
		Password: "weakpass", // no upper/special char
		Phone:    "12345",    // invalid
		Email:    "bademail", // invalid
		Age:      17,         // too young
	}

	// Validate the struct
	if err := validate.Struct(user); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on rule '%s'\n", e.Field(), e.Tag())
		}
		return
	}

	fmt.Println("All fields are valid!")
}

/*
OUTPUT:
$ go run main.go
Field 'Username' failed on rule 'notreserved'
Field 'Password' failed on rule 'strongpwd'
Field 'Phone' failed on rule 'phone'
Field 'Email' failed on rule 'email'
Field 'Age' failed on rule 'gte'
*/

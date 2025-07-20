package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Struct with built-in + custom tags
type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,notreserved"`
	Password string `json:"password" validate:"required,strongpwd"`
	Phone    string `json:"phone"    validate:"required,phone"`
	Email    string `json:"email"    validate:"required,email"`
	Age      int    `json:"age"      validate:"gte=18,lte=60"`
}

// i18n translations
var messages = map[string]map[string]string{
	"en": {
		"required":    "The field '{field}' is required.",
		"email":       "Invalid email format.",
		"notreserved": "Username '{field}' is not allowed.",
		"strongpwd":   "Password is too weak. Include upper, lower, and special characters.",
		"phone":       "Phone number must be a valid Indian mobile number.",
	},
	"hi": {
		"required":    "'{field}' फ़ील्ड आवश्यक है।",
		"email":       "ईमेल का प्रारूप अमान्य है।",
		"notreserved": "उपयोगकर्ता नाम '{field}' की अनुमति नहीं है।",
		"strongpwd":   "पासवर्ड बहुत कमजोर है। अपर/लोअर केस और विशेष वर्ण शामिल करें।",
		"phone":       "फ़ोन नंबर मान्य भारतीय मोबाइल नंबर होना चाहिए।",
	},
}

func main() {
	validate := validator.New()

	// Register custom rules
	validate.RegisterValidation("notreserved", func(fl validator.FieldLevel) bool {
		reserved := map[string]bool{"admin": true, "root": true}
		return !reserved[strings.ToLower(fl.Field().String())]
	})

	validate.RegisterValidation("strongpwd", func(fl validator.FieldLevel) bool {
		p := fl.Field().String()
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(p)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(p)
		hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_{}\[\]:;"'<>,.?/]`).MatchString(p)
		return len(p) >= 8 && hasUpper && hasLower && hasSpecial
	})

	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[6-9]\d{9}$`).MatchString(fl.Field().String())
	})

	// Simulated invalid input
	user := RegisterUserRequest{
		Username: "admin",        // Reserved
		Password: "pass123",      // Weak
		Phone:    "12345",        // Invalid format
		Email:    "invalidemail", // Invalid email
		Age:      17,             // Too young
	}

	// Choose language: "en" or "hi"
	lang := "hi"

	if err := validate.Struct(user); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			template := messages[lang][tag]
			if template == "" {
				template = fmt.Sprintf("Validation failed for '%s'", tag)
			}
			fmt.Printf("%s: %s\n", field, strings.ReplaceAll(template, "{field}", field))
		}
		return
	}

	fmt.Println("All fields are valid!")
}

/*
OUTPUT:
$ go run main.go
Username: उपयोगकर्ता नाम 'Username' की अनुमति नहीं है।
Password: पासवर्ड बहुत कमजोर है। अपर/लोअर केस और विशेष वर्ण शामिल करें।
Phone: फ़ोन नंबर मान्य भारतीय मोबाइल नंबर होना चाहिए।
Email: ईमेल का प्रारूप अमान्य है।
Age: Validation failed for 'gte'
*/

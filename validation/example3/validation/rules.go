package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var reservedUsernames = map[string]bool{
	"admin":  true,
	"root":   true,
	"system": true,
}

// RegisterCustomRules registers all custom validation functions
func RegisterCustomRules(v *validator.Validate) {
	// notreserved: username must not be in reserved list
	v.RegisterValidation("notreserved", func(fl validator.FieldLevel) bool {
		username := strings.ToLower(fl.Field().String())
		return !reservedUsernames[username]
	})

	// strongpwd: min 8 chars, must contain upper, lower, and special char
	v.RegisterValidation("strongpwd", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_{}\[\]:;"'<>,.?/]`).MatchString(password)
		return len(password) >= 8 && hasUpper && hasLower && hasSpecial
	})

	// phone: valid Indian mobile number
	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		return regexp.MustCompile(`^[6-9]\d{9}$`).MatchString(phone)
	})
}

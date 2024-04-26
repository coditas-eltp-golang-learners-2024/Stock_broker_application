package validations

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

// CustomValidator is a wrapper for validator.Validate with custom validation rules
type CustomValidator struct {
	Validator *validator.Validate
}

// NewCustomValidator initializes a new instance of the CustomValidator
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

// ValidateStruct validates a struct using custom validation rules
func (cv *CustomValidator) ValidateStruct(ctx context.Context, s interface{}) error {
	if err := cv.Validator.Struct(s); err != nil {
		// Format validation errors
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := FormatValidationErrors(validationErrors)

		// Prepare error message
		errorMessage := strings.Join(errorMessages, ", ")

		// Return error message
		return fmt.Errorf("validation failed: %s", errorMessage)
	}
	return nil
}

// FormatValidationErrors formats validation errors into a slice of strings
func FormatValidationErrors(validationErrors validator.ValidationErrors) []string {
	var errorMessages []string

	for _, err := range validationErrors {
		errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", err.Field(), err.Tag()))
	}

	return errorMessages
}

// ValidateEmail checks for a proper email format
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Regular expression for email format validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)

	return match
}

// ValidateStrongPassword checks for a strong password
func ValidateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Regular expression for strong password validation (at least 8 characters, at least one uppercase letter, one lowercase letter, one digit, and one special character)
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	match, _ := regexp.MatchString(passwordRegex, password)

	return match
}

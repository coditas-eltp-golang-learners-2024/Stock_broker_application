package validations
import (
	"context"
	"fmt"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"
	"github.com/go-playground/validator/v10"
)

var customValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	customValidator = validator.New()
	customValidator.RegisterValidation(constants.Password, ValidatePasswordStruct)
}

func GetCustomValidator(ctx context.Context) *validator.Validate {
	if customValidator == nil {
		NewCustomValidator(ctx)
	}
	return customValidator
}

// FormatValidationErrors formats validation errors into a user-friendly format
func FormatValidationErrors(ctx context.Context, validationErrors validator.ValidationErrors) ([]models.ErrorMessage, string) {
	var errorMessages []models.ErrorMessage
	var errorMessagesString []string

	// Iterate over each validation error and format it
	for _, err := range validationErrors {
		key := err.Field()
		message := err.Tag()

		// Add the error message to both structured error messages and string slice
		errorMessages = append(errorMessages, models.ErrorMessage{
			Key:          key,
			ErrorMessage: message,
		})
		errorMessagesString = append(errorMessagesString, fmt.Sprintf("%s: %s", key, message))
	}

	// Return both structured error messages and concatenated string
	return errorMessages, strings.Join(errorMessagesString, ", ")
}

// ValidatePasswordStruct is a custom validation function for password format.
func ValidatePasswordStruct(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	return isStrongPassword(input)
}

func isStrongPassword(password string) bool {
	// Define the criteria for a strong password
	minLength := 8
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecialChar := false

	// Check each character of the password
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUppercase = true
		case char >= 'a' && char <= 'z':
			hasLowercase = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case char >= 33 && char <= 47, char >= 58 && char <= 64, char >= 91 && char <= 96, char >= 123 && char <= 126:
			hasSpecialChar = true
		}
	}

	// Check if all criteria are met
	return len(password) >= minLength && hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}
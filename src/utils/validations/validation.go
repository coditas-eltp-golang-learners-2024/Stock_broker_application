package validations

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	genericConstants"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"
	"unicode"
)

var customValidator *validator.Validate

func NewCustomValidator(ctx context.Context) *validator.Validate {
	if customValidator == nil {
		customValidator = validator.New()
		customValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get(genericConstants.JsonConfig)
		})
		customValidator.RegisterValidation(genericConstants.CustomPasswordValidation, ValidatePasswordStruct)
	}
	return customValidator
}

func GetCustomValidator(ctx context.Context) *validator.Validate {
	if customValidator == nil {
		_ = NewCustomValidator(ctx)
	}
	return customValidator
}

// ValidatePasswordStruct is a custom validation function for password format.
func ValidatePasswordStruct(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	return isValidPassword(input)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}

var customErrorMap = map[string]string{
	"min":                genericConstants.MinValidationError,
	"required":           genericConstants.RequiredValidationError,
	"max":                genericConstants.MaxValidationError,
	"PasswordValidation": genericConstants.GenericPasswordValidationError,
}

// FormatValidationErrors formats validation errors into a user-friendly format
func FormatValidationErrors(ctx context.Context, validationErrors validator.ValidationErrors) ([]models.ErrorMessage, string) {
	var errorMessages []models.ErrorMessage
	var errorMessagesString []string

	// Iterate over each validation error and format it
	for _, err := range validationErrors {
		key := err.Field()
		message := err.Tag()
		errorParam := err.Param()
		var errorMessage string

		if errorParam != "" {
			errorMessage = fmt.Sprintf(customErrorMap[message], errorParam)
		} else {
			errorMessage = customErrorMap[message]
		}

		// Add the error message to both structured error messages and string slice
		errorMessages = append(errorMessages, models.ErrorMessage{
			Key:          key,
			ErrorMessage: errorMessage,
		})

		errorMessagesString = append(errorMessagesString, fmt.Sprintf(customErrorMap[message], 8))
	}

	// Return both structured error messages and concatenated string
	return errorMessages, strings.Join(errorMessagesString, ", ")
}

package validations

import (
	"context"
	"errors"
	"log"
	"reflect"
	genericConstants "stock_broker_application/src/constants"
	"unicode"

	"gopkg.in/go-playground/validator.v9"
)

var custValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	custValidator = validator.New()

	custValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(genericConstants.JsonConfig)
	})
	custValidator.RegisterValidation(genericConstants.PasswordValidation, ValidatePasswordStruct)

}

func GetCustomValidator(ctx context.Context) *validator.Validate {
	if custValidator == nil {
		NewCustomValidator(ctx)
	}
	return custValidator
}

// ValidatePasswordStruct is a custom validation function for password format.
func ValidatePasswordStruct(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	// Perform custom password format validation
	if err := validateCustomPasswordFormat(input); err != nil {
		// Custom validation failed
		log.Println("Custom password format validation error:", err)
		return false
	}

	return true
}

// validateCustomPasswordFormat is a custom function to validate password format.
func validateCustomPasswordFormat(input string) error {
	hasAlpha := false
	hasNumeric := false

	// Check each character in the password
	for _, char := range input {
		if unicode.IsLetter(char) {
			hasAlpha = true
		} else if unicode.IsDigit(char) {
			hasNumeric = true
		}

		// Early exit if both alphabetic and numeric characters are found
		if hasAlpha && hasNumeric {
			return nil
		}
	}
	// If password does not contain both alphabetic and numeric characters
	return errors.New(genericConstants.ErrValidatePassword)
}

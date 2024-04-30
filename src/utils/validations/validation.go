package validations

import (
	"context"
	"errors"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

var custValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	custValidator = validator.New()

	// Register custom validation functions
	custValidator.RegisterValidation("password", ValidatePasswordStruct)
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
	return validateCustomPasswordFormat(input) == nil
}

func validateCustomPasswordFormat(input string) error {
	strongPasswordRegex := `^[a-zA-Z0-9!@#$%^&*()_+=\-[\]{};:'",.<>/?]{8,}$`
	match, err := regexp.MatchString(strongPasswordRegex, input)
	if err != nil {
		return err
	}
	if match {
		return nil
	}
	return errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
}

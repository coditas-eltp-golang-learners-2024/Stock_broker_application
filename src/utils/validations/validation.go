package validations

import (
	"context"
	"errors"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"stock_broker_application/src/constants"
)

var custValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	custValidator = validator.New()
	custValidator.RegisterValidation(constants.Password, ValidatePasswordStruct)
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
	match, err := regexp.MatchString(constants.PasswordRegex, input)
	if err != nil {
		return err
	}
	if match {
		return nil
	}
	return errors.New(constants.ErrorValidatePassword)
}

package validations

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"stock_broker_application/src/constants"

	"gopkg.in/go-playground/validator.v9"
)

var custValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	custValidator = validator.New()

	custValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(constants.JsonConfig)
	})
	custValidator.RegisterValidation(constants.PasswordValidation, ValidatePasswordStruct)

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

	if err := validateCustomPasswordFormat(input); err != nil {
		return false
	}

	return true
}

func validateCustomPasswordFormat(input string) error {

	match, _ := regexp.MatchString(constants.PasswordRegex, input)

	if match {
		return nil
	}

	return errors.New(constants.ErrorValidatePassword)
}

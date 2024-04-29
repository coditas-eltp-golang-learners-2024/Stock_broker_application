package utils

import (
	"authentication/commons/constants"
	"github.com/go-playground/validator/v10"
	"regexp"
	genericConstants "stock_broker_application/src/constants"
)

func RegisterCustomValidations(validate *validator.Validate) error {
	if err := validate.RegisterValidation(genericConstants.CustomPasswordValidationTag, PasswordValidation); err != nil {
		return err
	}
	return nil
}

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	return regexp.MustCompile(constants.PasswordRegex).MatchString(password)
}

package commons

// Add your common functionalities here

import (
	"authentication/models"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func ValidateforgotPasswordStruct(validatePasswordResetRequest models.ForgotPasswordRequest) error {
	if err := validate.Struct(validatePasswordResetRequest); err != nil {
		return err
	}
	return nil
}

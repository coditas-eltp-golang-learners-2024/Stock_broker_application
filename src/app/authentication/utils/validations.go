package utils

import (
	"authentication/models"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func SignUpValidation(user *models.User) error {
	validate := validator.New()

	validate.RegisterValidation("passwordValidation", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		hasUpper := false
		hasDigit := false
		hasSpecial := false

		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsDigit(char):
				hasDigit = true
			case !unicode.IsLetter(char) && !unicode.IsDigit(char):
				hasSpecial = true
			}
		}

		return hasUpper && hasDigit && hasSpecial
	})

	err := validate.Struct(user)
	if err != nil {
		return err
	}

	return nil
}

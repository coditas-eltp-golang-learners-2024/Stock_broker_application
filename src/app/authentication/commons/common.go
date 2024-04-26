package commons

// Add your common functionalities here

import (
	"authentication/models"

	"gopkg.in/go-playground/validator.v9"
)

// Create a new instance of the validator to be used commonly for every other feature
var validate = validator.New()

// Function to perform validation
func ValidateForgetPasswordStruct(req models.ForgetPasswordRequest) error {
	if err := validate.Struct(req); err != nil {
		// Validation failed
		return err
	}
	return nil
}

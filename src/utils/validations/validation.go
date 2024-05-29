package validations

import (
	"context"
	"fmt"
	"reflect"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var custValidator *validator.Validate

func NewCustomValidator(ctx context.Context) {
	custValidator = validator.New()

	custValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(genericConstants.JsonConfig)
	})
	custValidator.RegisterValidation(genericConstants.CustomPasswordValidation, ValidatePasswordStruct)

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

	if result := isValid(input); result {
		return true
	}

	return false
}

var customErrorMap = map[string]string{
	"min":                genericConstants.MinValidationError,
	"required":           genericConstants.RequiredValidationError,
	"max":                genericConstants.MaxValidationError,
	"PasswordValidation": genericConstants.GenericPasswordValidationError,
}

var SliceErrors = make([]string, 0)

func isValid(newPassword string) bool {
	var (
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	for _, char := range newPassword {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if hasUpper && hasLower && hasNumber {
		return true
	}

	return false
}

// FormatValidationErrors formats validation errors into a user-friendly format
func FormatValidationErrors(ctx context.Context, validationErrors validator.ValidationErrors) []models.ErrorMessage {
	var errorMessages []models.ErrorMessage

	errorMessages = append(errorMessages, models.ErrorMessage{
		Key:          genericConstants.GenericJSONErrorMessage,
		ErrorMessage: genericConstants.ValidatorError})

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
	}

	return errorMessages
}

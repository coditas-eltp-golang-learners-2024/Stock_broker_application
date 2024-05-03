package validations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	// genericConstants "stock_broker_application/src/constants"
)
var CustomValidator *validator.Validate

// InitializeValidator initializes the custom validator
func InitializeValidator() {
	CustomValidator = validator.New()
	CustomValidator.RegisterValidation("validatePassword", ValidatePassword)
}

// ValidatePassword validates the password format
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Password criteria: at least one lowercase letter, one uppercase letter, one digit, and one special character
	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return regex.MatchString(password)
}

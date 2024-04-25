package constants

import "errors"

var (
	// Other errors
	ErrUserNotFound                      = errors.New("user not found")
	ErrInvalidEmail                      = errors.New("invalid email")
	ErrInvalidCredentials                = errors.New("invalid credentials")
	ErrEmptyField                        = errors.New("all required fields should be present")
	ErrValidationFailed                  = errors.New("validation failed")
	ErrEmailOrPasswordVerificationFailed = errors.New("email or password wrong")
	ErrShouldHaveNewPassword             = errors.New("please enter different password from old password")
	ErrInvalidPassword                   = errors.New("invalid password format")
)

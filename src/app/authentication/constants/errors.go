package constants

import "errors"

var (
	// Database errors
	ErrDatabaseConnection = errors.New("failed to connect to database")
	ErrDatabasePing       = errors.New("failed to ping database")
	ErrDatabaseQuery      = errors.New("database query failed")
	ErrDatabaseInsert     = errors.New("failed to insert into database")

	// Other errors
	ErrUserNotFound                      = errors.New("user not found")
	ErrInvalidCredentials                = errors.New("invalid credentials")
	ErrEmptyField                        = errors.New("all required fields should be present")
	ErrValidationFailed                  = errors.New("validation failed")
	ErrEmailOrPasswordVerificationFailed = errors.New("Email or password wrong")
	ErrShouldHaveNewPassword             = errors.New("Please enter different password from old password")
	ErrInvalidPassword                   = errors.New("invalid password format")
)

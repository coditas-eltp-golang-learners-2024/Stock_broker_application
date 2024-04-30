package constants

const (
	ErrorUserNotFound       = "user not found"
	ErrorInvalidPassword    = "invalid password"
	ErrorInvalidCredentials = "invalid credentials"
	ErrorInvalidEmailFormat = "invalid email address format"
	ErrorPasswordTooShort   = "password length must be at least 8 characters"
	ErrorFieldRequired      = "this field is required"
)

const (
	ErrorBadRequest                  = "Bad request"
	ErrorUnauthorized                = "Unauthorized"
	ErrorMessageAuthenticationFailed = "Authentication failed"
	ErrorAuthenticatingUserFormat    = "error authenticating user: %w"
)
const(
	InvalidPasswordFormatError = "password should contain both alphabetic and numeric characters."
	InvalidUserDataError       = "invalid user data."
	DefaultValidationError     = "error occurred while validating password format."
)

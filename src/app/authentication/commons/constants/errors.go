package constants

const (
	ErrUserNotFound       = "user not found"
	ErrInvalidPassword    = "invalid password"
	ErrInvalidCredentials = "invalid credentials"
	ErrInvalidEmailFormat = "invalid email address format"
	ErrPasswordTooShort   = "password length must be at least 8 characters"
	ErrFieldRequired      = "this field is required"
)

const (
	ErrorBadRequest                  = "Bad request"
	ErrorUnauthorized                = "Unauthorized"
	ErrorMessageAuthenticationFailed = "Authentication failed"
	SuccessMessageSignIn             = "User authenticated successfully"
	ErrorAuthenticatingUserFormat    = "error authenticating user: %w"
)

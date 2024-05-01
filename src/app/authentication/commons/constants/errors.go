package constants

const (
	ErrorInvalidCredentials          = "invalid credentials"
	ErrorBadRequest                  = "Bad request"
	ErrorUnauthorized                = "Unauthorized"
	ErrorMessageAuthenticationFailed = "Authentication failed"
	ErrorAuthenticatingUser          = "error authenticating user: %w"
	InvalidPasswordFormatError       = "password should contain both alphabetic and numeric characters."
	ErrorGenerateAndSaveOTP          = "failed to generate and save OTP"
)

package constants

const (
	ErrorInvalidCredentials          = "invalid credentials"
	ErrorBadRequest                  = "bad request"
	ErrorUnauthorized                = "unauthorized"
	ErrorMessageAuthenticationFailed = "authentication failed"
	ErrorAuthenticatingUser          = "error authenticating user: %w"
	ErrorGenerateAndSaveOTP          = "failed to generate and save OTP"
)

// Errors related to Forgot-Password Validations
const (
	ErrorInvalidUserData        = "invalid user data."
	ErrorWatchlistNotFound      = "watchlist not found for the given user ID and name"
)

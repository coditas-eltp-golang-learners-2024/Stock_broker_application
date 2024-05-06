package constants

// Add your error constants here

var (
	ErrUserExists = "User already exists"
)

// Errors related to Forgot-Password Validations
const (
	ErrorOtpVerification             = "otp verification failed"
	ErrorGenToken                    = "failed to generate token"
	ErrorRequiredUsername            = "username is required"
	ErrorInvalidOtpFormat            = "invalid otp format"
	ErrorInvalidOTP                  = "invalid otp"
	ErrorInvalidUsername             = "invalid username"
	ErrorCheckingOtp                 = "failed to check otp"
	ErrorCheckUserExists             = "failed to check user exists"
	ErrorUpdateUserToken             = "failed to update user token"
	ErrorInvalidRequest              = "invalid request"
	ErrorInvalidUserData             = "invalid user data."
	ErrorInvalidCredentials          = "invalid credentials"
	ErrorBadRequest                  = "bad request"
	ErrorUnauthorized                = "unauthorized"
	ErrorMessageAuthenticationFailed = "authentication failed"
	ErrorAuthenticatingUser          = "error authenticating user: %w"
	ErrorGenerateAndSaveOTP          = "failed to generate and save OTP"
)

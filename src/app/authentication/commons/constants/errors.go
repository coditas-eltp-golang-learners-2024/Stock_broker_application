package constants

// authentication errors
const (
	ErrorUserExists                  = "user already exists"
	ErrorInvalidUserIDOrPassword     = "invalid UserID or password"
	ErrorFailedToSetNewPassword      = "failed to set a new password"
	ErrorOtpVerification             = "otp verification failed"
	ErrorGenToken                    = "failed to generate token"
	ErrorRequiredUserID              = "userID is required"
	ErrorInvalidOtpFormat            = "invalid otp format"
	ErrorInvalidOTP                  = "invalid otp"
	ErrorInvalidUserID               = "invalid userid"
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

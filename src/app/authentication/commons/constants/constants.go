package constants

// authentication NEST API URL Keys
const (
	ServiceName          = "authentication"
	PortDefaultValue     = 8080
	DatabaseYamlFilePath = "resources"
	DatabasePort         = "3306"
)

// Authentication success message
const (
	ValidateOTPSuccessMessage    = "OTP validated successfully"
	ChangePasswordSuccessMessage = "Password changed successfully"
	ForgotPasswordSuccessMessage = "Password updated successfully"
	SignInSuccessMessage         = "user authenticated successfully"
)

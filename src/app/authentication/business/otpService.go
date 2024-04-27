package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"stock_broker_application/src/utils/authorization"
)

type OTPService struct {
	UserRepository repositories.CustomerRepository
}

func NewOTPService(userRepository repositories.CustomerRepository) *OTPService {
	return &OTPService{
		UserRepository: userRepository,
	}
}
func (otpService *OTPService) OtpVerification(otpData models.OTPValidationRequest) error {
	if !otpService.UserRepository.CheckOtp(otpData.Email, otpData.OTP) {
		return constants.ErrOtpVerification
	}
	return nil
}

func (otpService *OTPService) GenerateAndStoreToken(email string) (string, error) {
	// Generate JWT token
	token, err := authorization.GenerateJWTToken(email)
	if err != nil {
		return "", err
	}

	// Store token in the database
	if err := otpService.UserRepository.UpdateUserToken(email, token); err != nil {
		return "", err
	}

	return token, nil
}

package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
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
func (otpService *OTPService) OtpVerification(otpData models.Users) error {
	if otpService.UserRepository.CheckOtp(otpData.Email, otpData.OTP) {
		return errors.New(constants.ErrorOtpVerification)
	}
	return nil
}

func (otpService *OTPService) GenerateAndStoreToken(email string) (string, error) {
	token, err := authorization.GenerateJWTToken(email)
	if err != nil {
		return "", err
	}

	if err := otpService.UserRepository.UpdateUserToken(email, token); err != nil {
		return "", err
	}

	return token, nil
}

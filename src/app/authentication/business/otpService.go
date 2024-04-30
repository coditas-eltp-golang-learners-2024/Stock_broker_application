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
	if otpService.UserRepository.CheckOtp(otpData.UserName, otpData.OTP) {
		return errors.New(constants.ErrorOtpVerification)
	}
	return nil
}

func (otpService *OTPService) GenerateAndStoreToken(username string) (string, error) {
	token, err := authorization.GenerateJWTToken(username)
	if err != nil {
		return "", err
	}

	if err := otpService.UserRepository.UpdateUserToken(username, token); err != nil {
		return "", err
	}

	return token, nil
}

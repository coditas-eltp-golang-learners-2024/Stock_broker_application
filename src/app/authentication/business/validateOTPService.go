package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	genericModel "stock_broker_application/src/models"
	"stock_broker_application/src/utils/authorization"
)

type OTPService struct {
	UserRepository repositories.UserRepository
}

func NewOTPService(userRepository repositories.UserRepository) *OTPService {
	return &OTPService{
		UserRepository: userRepository,
	}
}

func (service *OTPService) OtpVerification(otpData models.ValidateOTPRequest) error {
	if otpData.UserName == "" {
		return errors.New(constants.ErrorRequiredUsername)
	}
	if otpData.OTP < 1000 || otpData.OTP > 9999 {
		return errors.New(constants.ErrorInvalidOtpFormat)
	}

	userExists, err := service.UserRepository.CheckUserExists(otpData.UserName)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New(constants.ErrorInvalidUsername)
	}

	isValid, err := service.UserRepository.CheckOtp(otpData.UserName, otpData.OTP)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New(constants.ErrorInvalidOtpFormat) // This line was updated
	}
	return nil
}

func (service *OTPService) GenerateAndStoreToken(tokenData genericModel.TokenData, username string) (string, error) {
	token, err := authorization.GenerateJWTToken(tokenData)
	if err != nil {
		return "", err
	}

	if err := service.UserRepository.UpdateUserToken(username, token); err != nil {
		return "", err
	}

	return token, nil
}

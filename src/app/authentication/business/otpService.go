package business

import (
	"authentication/repositories"
	"math/rand"
)

type OTPService struct {
	userRepository repositories.UserRepository
}

func NewOTPService(userRepository repositories.UserRepository) *OTPService {
	return &OTPService{
		userRepository: userRepository,
	}
}
func (otpservice *OTPService) GenerateAndSaveOTP(username string) error {
	otp := rand.Intn(8999) + 1000
	if err := otpservice.userRepository.UpdateOTPAndCreationTime(username, otp); err != nil {
		return err
	}
	return nil
}

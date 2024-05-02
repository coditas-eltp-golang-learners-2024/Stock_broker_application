package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"fmt"
	"math/rand"
)

// SignInService handles the sign-in logic
type SignInService struct {
	userRepository repositories.SignInRepository
}

// NewSignInService creates a new SignInService with the provided repository
func NewSignInService(userRepository repositories.SignInRepository) *SignInService {
	return &SignInService{
		userRepository: userRepository,
	}
}

// SignIn verifies the sign-in credentials
func (service *SignInService) SignIn(signInRequest models.SignInRequest) error {
	user, err := service.userRepository.AuthenticateUser(signInRequest.UserName, signInRequest.Password)
	if err != nil {
		return fmt.Errorf(constants.ErrorAuthenticatingUser, err)
	}
	if !user {
		return errors.New(constants.ErrorInvalidCredentials)
	}
	return nil
}
func (service *SignInService) GenerateAndSaveOTP(username string) error {
	otp := rand.Intn(8999) + 1000
	if err := service.userRepository.UpdateOTPAndCreationTime(username, otp); err != nil {
		return err
	}
	return nil
}


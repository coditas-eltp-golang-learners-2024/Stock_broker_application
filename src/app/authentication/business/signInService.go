package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"fmt"
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
	user := service.userRepository.GetUserByUsername(signInRequest.UserName)
	if user == nil {
		return errors.New(constants.ErrorUserNotFound)
	}
	ok, err := AuthenticateUser(user, signInRequest.Password)
	if err != nil {
		return fmt.Errorf(constants.ErrorAuthenticatingUserFormat, err)
	} else if !ok {
		return errors.New(constants.ErrorInvalidCredentials)
	}
	return nil
}

// AuthenticateUser verifies the user's password
func AuthenticateUser(user *models.SignInRequest, password string) (bool, error) {
	if user == nil {
		return false, errors.New(constants.ErrorInvalidPassword)
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

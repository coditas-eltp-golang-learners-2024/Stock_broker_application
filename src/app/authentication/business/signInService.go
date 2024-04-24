package business

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repositories"
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
func (s *SignInService) SignIn(signInRequest models.SignInRequest) error {
	// Retrieve user from the database based on the provided email
	user := s.userRepository.GetUserByEmail(signInRequest.Email)

	// Check if user exists
	if user == nil {
		return constants.ErrUserNotFound
	}

	// Verify the password
	if ok, err := AuthenticateUser(user, signInRequest.Password); err != nil {
		return err
	} else if !ok {
		return constants.ErrInvalidCredentials
	}

	// Sign-in successful
	return nil
}

// AuthenticateUser verifies the user's password
func AuthenticateUser(user *models.SignInRequest, password string) (bool, error) {
	if user == nil {
		return false, constants.ErrInvalidPassword
	}

	// Compare the provided password with the password stored in the database
	if user.Password != password {
		// Passwords do not match
		return false, nil
	}

	// Passwords match
	return true, nil
}

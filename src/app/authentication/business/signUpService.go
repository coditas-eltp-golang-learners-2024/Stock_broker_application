package business

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
)

type SignUpService struct {
	UserSignUpRepository repositories.UserSignUpRepository
}

func NewSignUpService(userSignUpRepository repositories.UserSignUpRepository) *SignUpService {
	return &SignUpService{
		UserSignUpRepository: userSignUpRepository,
	}
}

func (b *SignUpService) SignUp(user *models.User) error {

	count, err := b.UserSignUpRepository.CheckUserExists(user)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constants.ErrUserExists)
	}
	result := b.UserSignUpRepository.InsertUserIntoDB(user)
	if result != nil {
		return result
	}
	return nil
}

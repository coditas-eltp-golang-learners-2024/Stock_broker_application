package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"stock_broker_application/src/utils/postgres"
)

type SignUpService struct {
	UserSignUpRepository repositories.UserSignUpRepository
}

func NewSignUpService(userSignUpRepository repositories.UserSignUpRepository) *SignUpService {
	return &SignUpService{
		UserSignUpRepository: userSignUpRepository,
	}
}

func (service *SignUpService) SignUp(user *models.UserSignUp) error {
	client := postgres.GetPostGresClient()
	count, err := service.UserSignUpRepository.CheckUserExists(client.GormDb, user)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constants.ErrUserExists)
	}
	result := service.UserSignUpRepository.InsertUserIntoDB(client.GormDb, user)
	if result != nil {
		return result
	}
	return nil
}

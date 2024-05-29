package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	genericModels "stock_broker_application/src/models"
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
	dbUser := genericModels.Users{
		UserName:    user.UserName,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		PanCard:     user.PanCard,
		Password:    user.Password,
	}
	count, err := service.UserSignUpRepository.CheckUserExists(client.GormDb, &dbUser)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New(constants.ErrorUserExists)
	}

	result := service.UserSignUpRepository.InsertUserIntoDB(client.GormDb, &dbUser)
	if result != nil {
		return result
	}
	return nil
}

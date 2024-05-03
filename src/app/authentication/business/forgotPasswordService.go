package business

import (
	"authentication/models"
	"authentication/repositories"
)

type NewforgotPasswordService interface {
	UpdatePassword(credentials models.ForgotPasswordRequest) error
}

type userDataService struct {
	userCredentialsInterface repositories.ForgotPasswordRepository
}

func NewUsersService(userData repositories.ForgotPasswordRepository) NewforgotPasswordService {
	return &userDataService{
		userCredentialsInterface: userData,
	}
}

func (service *userDataService) UpdatePassword(credentials models.ForgotPasswordRequest) error {

	return service.userCredentialsInterface.VerifyAndUpdatePassword(credentials.Email, credentials.PanCardNumber, credentials.Password)
}

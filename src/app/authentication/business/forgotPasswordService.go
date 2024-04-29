package business

import (
	"authentication/models"
	"authentication/repositories"
)

type NewforgotPasswordService interface {
	UpdatePassword(credentials models.ForgotPasswordRequest) error
}

type credentialsService struct {
	userCredentialsInterface repositories.ForgotPasswordRequest
}

func NewCredentialsService(userData repositories.ForgotPasswordRequest) NewforgotPasswordService {
	return &credentialsService{
		userCredentialsInterface: userData,
	}
}

func (service *credentialsService) UpdatePassword(credentials models.ForgotPasswordRequest) error {

	return service.userCredentialsInterface.VerifyCredentialsAndUpdateOTP(credentials.Email, credentials.PanCardNumber, credentials.Password)
}

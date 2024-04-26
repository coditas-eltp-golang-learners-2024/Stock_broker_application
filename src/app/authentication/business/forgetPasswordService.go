package business

import (
	"authentication/models"
	"authentication/repositories"
)

type ForgetPasswordService interface {
	UpdatePassword(credentials models.ForgetPasswordRequest) error
}

type credentialsService struct {
	repo repositories.ForgetPasswordRequest
}

func NewCredentialsService(repo repositories.ForgetPasswordRequest) ForgetPasswordService {
	return &credentialsService{
		repo: repo,
	}
}

func (service *credentialsService) UpdatePassword(credentials models.ForgetPasswordRequest) error {

	return service.repo.VerifyCredentialsAndUpdateOTP(credentials.Email, credentials.PanCardNumber, credentials.Password)
}

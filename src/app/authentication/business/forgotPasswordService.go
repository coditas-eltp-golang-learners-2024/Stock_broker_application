package business

import (
	"authentication/models"
	"authentication/repositories"
)

type NewforgotPasswordService interface {
	UpdatePassword(credentials models.ForgotPasswordRequest) error
}

type forgotPasswordSercvice struct {
	userCredentialsInterface repositories.ForgotPasswordRepository
}

func NewUsersService(userData repositories.ForgotPasswordRepository) NewforgotPasswordService {
	return &forgotPasswordSercvice{
		userCredentialsInterface: userData,
	}
}

func (service *forgotPasswordSercvice) UpdatePassword(userModel models.ForgotPasswordRequest) error {

	return service.userCredentialsInterface.VerifyAndUpdatePassword(userModel.Email, userModel.PanCardNumber, userModel.NewPassword)
}

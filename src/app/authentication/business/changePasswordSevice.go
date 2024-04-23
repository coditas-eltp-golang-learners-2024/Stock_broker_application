package business

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repositories"
)

type PasswordResetter struct {
	PasswordResetter repositories.AuthenticationProvider
}

func NewRestPasswordService(restPswd repositories.AuthenticationProvider) *PasswordResetter {
	return &PasswordResetter{
		PasswordResetter: restPswd,
	}
}

func (service *PasswordResetter) ResetPassword(request models.ChangePassword) error {
	if !service.PasswordResetter.CheckEmailAndPassword(request.Email, request.OldPassword) {
		return constants.ErrInvalidEmailOrPassword
	}
	if !service.PasswordResetter.SetNewPassword(request.Email, request.NewPassword) {
		return constants.ErrFailedToSetNewPassword
	}
	return nil
}

package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"authentication/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type PasswordResetter struct {
	PasswordResetter repositories.AuthenticationProvider
}

func NewRestPasswordService(restPswd repositories.AuthenticationProvider) *PasswordResetter {
	return &PasswordResetter{
		PasswordResetter: restPswd,
	}
}

func (service *PasswordResetter) ResetPassword(request models.ChangePassword, ctx *gin.Context) error {
	// Retrieve the value associated with the key "email" from ctx
	email := ctx.Value(constants.EmailId).(string)
	if !service.PasswordResetter.CheckEmailAndPassword(email, request.OldPassword) {
		return errors.New(constants.ErrInvalidEmailOrPassword)
	}
	utils.ValidatePassword(request.OldPassword)
	if !service.PasswordResetter.SetNewPassword(email, request.NewPassword) {
		return errors.New(constants.ErrFailedToSetNewPassword)
	}
	return nil
}

package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"github.com/gin-gonic/gin"
)

type ChangePassword struct {
	ChangePasswordRepository repositories.ChangePasswordRepositor
}

func NewRestPasswordService(changePasswordInstance repositories.ChangePasswordRepositor) *ChangePassword {
	return &ChangePassword{
		ChangePasswordRepository: changePasswordInstance,
	}
}

func (service *ChangePassword) ChangePasswordService(request models.ChangePassword, ctx *gin.Context) error {
	username := ctx.Value(constants.UserName).(string)
	userCheckQuery := map[string]interface{}{
		constants.UserName: username,
		constants.Password: request.OldPassword,
	}
	if !service.ChangePasswordRepository.CheckUsernameAndPassword(userCheckQuery) {
		return errors.New(constants.ErrorInvalidUsernameOrPassword)
	}
	passwordChangeQuery := map[string]interface{}{
		constants.UserName: username,
		constants.Password: request.NewPassword,
	}
	if !service.ChangePasswordRepository.SetNewPassword(passwordChangeQuery) {
		return errors.New(constants.ErrorFailedToSetNewPassword)
	}
	return nil
}

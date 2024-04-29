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
	PasswordResetterRepository repositories.AuthenticationProvider
}

func NewRestPasswordService(restPswd repositories.AuthenticationProvider) *PasswordResetter {
	return &PasswordResetter{
		PasswordResetterRepository: restPswd,
	}
}

func (service *PasswordResetter) ResetPassword(request models.ChangePassword, ctx *gin.Context) error {
	email := ctx.Value(constants.EmailId).(string)
	condition := map[string]interface{}{
		constants.EmailId:  email,
		constants.Password: request.OldPassword,
	}
	if !service.PasswordResetterRepository.CheckEmailAndPassword(condition) {
		return errors.New(constants.ErrorInvalidEmailOrPassword)
	}
	err := utils.ValidatePassword(request.NewPassword)
	if err != nil {
		return err
	}
	PasswordUpdateSQLCondition := map[string]interface{}{
		constants.EmailId:  email,
		constants.Password: request.NewPassword,
	}
	if !service.PasswordResetterRepository.SetNewPassword(PasswordUpdateSQLCondition) {
		return errors.New(constants.ErrorFailedToSetNewPassword)
	}
	return nil
}

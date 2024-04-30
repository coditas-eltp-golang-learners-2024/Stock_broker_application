package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"stock_broker_application/src/utils/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	username := ctx.Value(constants.UserName).(string)
	condition := map[string]interface{}{
		constants.UserName:  username,
		constants.Password: request.OldPassword,
	}
	if !service.PasswordResetterRepository.CheckEmailAndPassword(condition) {
		return errors.New(constants.ErrorInvalidEmailOrPassword)
	}
	validate := validator.New()
	if err := validations.RegisterCustomValidations(validate); err != nil {
		return errors.New(err.Error())
	}
	if err := validate.Struct(request); err != nil {
		return errors.New(err.Error())
	}
	PasswordUpdateSQLCondition := map[string]interface{}{
		constants.UserName:  username,
		constants.Password: request.NewPassword,
	}
	if !service.PasswordResetterRepository.SetNewPassword(PasswordUpdateSQLCondition) {
		return errors.New(constants.ErrorFailedToSetNewPassword)
	}
	return nil
}

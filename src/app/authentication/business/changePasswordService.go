package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
)

type ChangePassword struct {
	ChangePasswordRepository repositories.ChangePasswordRepositor
}

func NewChangePasswordService(changePasswordInstance repositories.ChangePasswordRepositor) *ChangePassword {
	return &ChangePassword{
		ChangePasswordRepository: changePasswordInstance,
	}
}

func (service *ChangePassword) ChangePasswordService(request models.ChangePassword, ctx *gin.Context) error {
	username := ctx.Value(genericConstants.Username).(string)
	userCheckQuery := map[string]interface{}{
		genericConstants.Username: username,
		genericConstants.Password: request.OldPassword,
	}
	if !service.ChangePasswordRepository.CheckUsernameAndPassword(userCheckQuery) {
		return errors.New(constants.ErrorInvalidUsernameOrPassword)
	}
	passwordChangeQuery := map[string]interface{}{
		genericConstants.Username: username,
		genericConstants.Password: request.NewPassword,
	}
	if !service.ChangePasswordRepository.SetNewPassword(passwordChangeQuery) {
		return errors.New(constants.ErrorFailedToSetNewPassword)
	}
	return nil
}

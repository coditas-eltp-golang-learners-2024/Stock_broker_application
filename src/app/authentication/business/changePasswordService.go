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
	userID := ctx.Value(genericConstants.Id).(string)
	userCheckQuery := map[string]interface{}{
		genericConstants.Id:       userID,
		genericConstants.Password: request.OldPassword,
	}
	if !service.ChangePasswordRepository.CheckUserIDAndPassword(userCheckQuery) {
		return errors.New(constants.ErrorInvalidUserIDOrPassword)
	}
	passwordChangeQuery := map[string]interface{}{
		genericConstants.Id:       userID,
		genericConstants.Password: request.NewPassword,
	}
	if !service.ChangePasswordRepository.SetNewPassword(passwordChangeQuery) {
		return errors.New(constants.ErrorFailedToSetNewPassword)
	}
	return nil
}

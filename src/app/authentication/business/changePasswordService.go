package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repositories"
	"errors"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
)

type ChangePasswordService struct {
	ChangePasswordRepository repositories.ChangePasswordRepository
}

func NewChangePasswordService(changePasswordInstance repositories.ChangePasswordRepository) *ChangePasswordService {
	return &ChangePasswordService{
		ChangePasswordRepository: changePasswordInstance,
	}
}

func (service *ChangePasswordService) ChangePasswordService(request models.ChangePassword, ctx *gin.Context) error {
	userID := ctx.Value(genericConstants.Id).(uint16)
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

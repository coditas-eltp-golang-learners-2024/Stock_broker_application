package handler

import (
	"authentication/business"
	genericConstants "authentication/commons/constants"
	"authentication/models"
	"fmt"
	"net/http"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
)

type ControllerChangePassword struct {
	ChangePasswordController *business.ChangePassword
}

func NewChangePasswordController(service *business.ChangePassword) *ControllerChangePassword {
	return &ControllerChangePassword{ChangePasswordController: service}
}

// @Summary Change Password
// @Description Change a user's password
// @Tags Authentication
// @Accept json
// @Produce json
// @Security JWT
// @Param request body models.ChangePassword true "Change Password Request"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/change-password [patch]
func HandleChangePassword(service *ControllerChangePassword) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var changeRequest models.ChangePassword
		if err := ctx.BindJSON(&changeRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.GenericErrorMessage: err.Error()})
			return
		}
		// Validate the change password request
		err := validations.CustomValidator.Struct(changeRequest)
		if err != nil {
			fmt.Println("Validation error:", err)
			return
		}

		if err := service.ChangePasswordController.ChangePasswordService(changeRequest, ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.GenericErrorMessage: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{genericConstants.StatusKey: genericConstants.ChangePasswordSuccessMessage})
	}
}

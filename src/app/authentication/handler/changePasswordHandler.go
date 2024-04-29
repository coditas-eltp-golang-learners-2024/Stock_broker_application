package handler

import (
	"authentication/business"
	genericConstants "authentication/commons/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/src/constants"
)

// @Summary Change Password
// @Description Change a user's password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.ChangePassword true "Change Password Request"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /changepassword [patch]
func HandleChangePassword(service *business.PasswordResetter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var changeRequest models.ChangePassword
		if err := ctx.BindJSON(&changeRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.GenericErrorMessage: err.Error()})
			return
		}
		if err := service.ResetPassword(changeRequest, ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.GenericErrorMessage: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{genericConstants.StatusKey: genericConstants.PasswordChangedSuccessMessage})
	}
}

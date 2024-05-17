package handler

import (
	"authentication/business"
	serviceConstants "authentication/commons/constants"
	"authentication/models"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChangePasswordController struct {
	ChangePasswordController *business.ChangePasswordService
}

func NewChangePasswordController(service *business.ChangePasswordService) *ChangePasswordController {
	return &ChangePasswordController{ChangePasswordController: service}
}

// @Summary Change Password
// @Description Change a user's password
// @Tags Change Password
// @Accept json
// @Produce json
// @Security JWT
// @Param request body models.ChangePassword true "Change Password Request"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/auth/change-password [patch]
func HandleChangePassword(service *ChangePasswordController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var changeRequest models.ChangePassword
		if err := ctx.BindJSON(&changeRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}
		if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(changeRequest); err != nil {
			validationErrors := validations.FormatValidationErrors(ctx.Request.Context(), err.(validator.ValidationErrors))
			ctx.JSON(http.StatusBadRequest, gin.H{
				genericConstants.GenericJSONErrorMessage: genericConstants.ValidatorError,
				genericConstants.GenericValidationError:  validationErrors,
			})
			return
		}
		if err := service.ChangePasswordController.ChangePasswordService(changeRequest, ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{genericConstants.BFFResponseSuccessMessage: serviceConstants.ChangePasswordSuccessMessage})
	}
}

package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateCredentialsHandler updates the user's credentials.
// @Summary Update user credentials
// @Description Updates user credentials based on the provided request.
// @ID update-credentials
// @Accept  json
// @Produce  json
// @Param request body models.ForgotPasswordRequest true "Forgot Password Request JSON"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /forgot-password [post]
func UpdateCredentialsHandler(service business.NewforgotPasswordService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request models.ForgotPasswordRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}

		// Perform validation using the initialized validator
		if err := validations.GetCustomValidator(context.Request.Context()).Struct(request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: constants.InvalidPasswordFormatError})
			return
		}

		// var request dbModels.ForgotPasswordRequest
		err := service.UpdatePassword(request) // Use 'service' instead of 'business'
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.InvalidUserDataError})
			return
		}
		context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.PasswordUpdated})
	}
}

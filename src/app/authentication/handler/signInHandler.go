package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"
)

// SignInHandler handles the sign-in request
// @Summary Handle sign-in request
// @Description Handle sign-in request and authenticate the user
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {object} string "User authenticated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /v1/authentication/signin [post]
func SignInHandler(userService *business.SignInService, otpService *business.OTPService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signInRequest models.SignInRequest

		if err := context.ShouldBindJSON(&signInRequest); err != nil {
			context.JSON(http.StatusBadRequest, constants.ErrorBadRequest)
			return
		}
		if err := validations.GetCustomValidator(context.Request.Context()).Struct(signInRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}
		if err := userService.SignIn(signInRequest); err != nil {
			context.JSON(http.StatusUnauthorized, constants.ErrorMessageAuthenticationFailed)
			return
		}

		context.JSON(http.StatusOK, constants.SuccessMessageSignIn)

		if err := otpService.GenerateAndSaveOTP(signInRequest.UserName); err != nil {
			context.JSON(http.StatusInternalServerError, constants.ErrorGenerateAndSaveOTP)
			return
		}

	}
}

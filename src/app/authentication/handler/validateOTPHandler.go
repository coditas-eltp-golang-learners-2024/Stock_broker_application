package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
)

// NewValidateOTPHandler handles the OTP validation request
// @Summary Validate OTP
// @Description Validates the OTP for a user
// @Tags OTP
// @Accept json
// @Produce json
// @Param otpRequest body models.OTPValidationRequest true "OTP Request"
// @Success 200 {string} string "OTP validated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "OTP is expired or invalid"
// @Router /validateOTP [post]
func NewValidateOTPHandler(otpService *business.OTPService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var otpValidationRequest models.ValidateOTPRequest

		// Bind JSON request body to OTPValidationRequest struct
		if err := context.ShouldBindJSON(&otpValidationRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}

		// Call OTPVerification method to validate OTP
		if err := otpService.OtpVerification(otpValidationRequest); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
			return
		}

		// Generate and store JWT token using OTPService
		token, err := otpService.GenerateAndStoreToken(otpValidationRequest.UserName)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.ErrorGenToken})
			return
		}

		// OTP validation success
		context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.SuccessOTPValidation, genericConstants.GenericTokenMessage: token})
	}
}

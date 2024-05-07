package handler

import (
	"authentication/commons/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
)

type OTPService interface {
	OtpVerification(otpData models.ValidateOTPRequest) error
	GenerateAndStoreToken(tokenData genericModel.TokenData, username string) (string, error)
}

type OTPValidationController struct {
	Service OTPService
}

func NewOTPValidationController(service OTPService) *OTPValidationController {
	return &OTPValidationController{
		Service: service,
	}
}

// NewValidateOTPHandler handles the OTP validation request
// @Summary Validate OTP
// @Description Validates the OTP for a user
// @Tags OTP
// @Accept json
// @Produce json
// @Param otpRequest body models.ValidateOTPRequest true "OTP Request"
// @Success 200 {string} string "OTP validated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "OTP is expired or invalid"
// @Router /validateOTP [post]
func (controller *OTPValidationController) HandleValidateOTP(context *gin.Context) {
	var otpValidationRequest models.ValidateOTPRequest
	if err := context.ShouldBindJSON(&otpValidationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: constants.ErrorInvalidRequest})
		return
	}
	if err := controller.Service.OtpVerification(otpValidationRequest); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	tokenData := genericModel.TokenData{
		UserId: otpValidationRequest.UserID,
	}
	token, err := controller.Service.GenerateAndStoreToken(tokenData, otpValidationRequest.UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.ErrorGenToken})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.ValidateOTPSuccessMessage, genericConstants.GenericTokenMessage: token})
}

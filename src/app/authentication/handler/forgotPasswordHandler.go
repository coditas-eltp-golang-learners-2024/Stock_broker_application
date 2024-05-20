package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"encoding/json"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
	"stock_broker_application/src/utils"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewForgetPasswordController interface {
	HandleForgotPassword(ctx *gin.Context)
}

type forgotPasswordController struct {
	service business.NewforgotPasswordService
}

func NewUsersController(service business.NewforgotPasswordService) NewForgetPasswordController {
	return &forgotPasswordController{
		service: service,
	}
}

// UpdateCredentialsHandler updates the user's credentials.
// @Summary Update user credentials
// @Description Updates user credentials based on the provided request.
// @Tags Forgot Password
// @ID update-credentials
// @Accept  json
// @Produce  json
// @Param request body models.ForgotPasswordRequest true "Forgot Password Request JSON"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/auth/forgot-password [post]
func (controller *forgotPasswordController) HandleForgotPassword(context *gin.Context) {
	var request models.ForgotPasswordRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		utils.SendBadRequest(context, []genericModel.ErrorMessage{errorMsgs})
		return
	}
	if err := validations.GetCustomValidator(context.Request.Context()).Struct(request); err != nil {
		validationErrors := validations.FormatValidationErrors(context.Request.Context(), err.(validator.ValidationErrors))
		context.JSON(http.StatusBadRequest, gin.H{
			genericConstants.GenericJSONErrorMessage: genericConstants.ValidatorError,
			genericConstants.GenericValidationError:  validationErrors,
		})
		return
	}
	err := controller.service.UpdatePassword(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.ErrorInvalidUserData})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.ForgotPasswordSuccessMessage})
	utils.SendStatusOkSuccess(context, constants.ForgotPasswordSuccessMessage)
}

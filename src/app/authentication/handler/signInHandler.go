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

type signInController struct {
	service *business.SignInService
}

func NewSignInController(service *business.SignInService) *signInController {
	return &signInController{
		service: service,
	}
}

// SignInHandler handles the sign-in request
// @Summary Handle sign-in request
// @Description Handle sign-in request and authenticate the user
// @Tags SignIN
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {object} string "User authenticated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /v1/auth/signin [post]
func (controller *signInController) HandleSignIn(context *gin.Context) {
	var signInRequest models.SignInRequest

	if err := context.ShouldBindJSON(&signInRequest); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		utils.SendBadRequest(context, []genericModel.ErrorMessage{errorMsgs})
		return
	}
	if err := validations.GetCustomValidator(context.Request.Context()).Struct(signInRequest); err != nil {
		validationErrors := validations.FormatValidationErrors(context.Request.Context(), err.(validator.ValidationErrors))
		context.JSON(http.StatusBadRequest, gin.H{
			genericConstants.GenericJSONErrorMessage: genericConstants.ValidatorError,
			genericConstants.GenericValidationError:  validationErrors,
		})

		return
	}
	if err := controller.service.SignIn(signInRequest); err != nil {
		context.JSON(http.StatusUnauthorized, constants.ErrorMessageAuthenticationFailed)
		return
	}

	context.JSON(http.StatusOK, constants.SignInSuccessMessage)

	if err := controller.service.GenerateAndSaveOTP(signInRequest.UserName); err != nil {
		context.JSON(http.StatusInternalServerError, constants.ErrorGenerateAndSaveOTP)
		return
	}
}

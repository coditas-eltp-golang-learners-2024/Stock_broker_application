package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
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
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {object} string "User authenticated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /v1/signin [post]
func (controller *signInController) HandleSignIn(context *gin.Context) {
	var signInRequest models.SignInRequest

	if err := context.ShouldBindJSON(&signInRequest); err != nil {
		context.JSON(http.StatusBadRequest, constants.ErrorBadRequest)
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

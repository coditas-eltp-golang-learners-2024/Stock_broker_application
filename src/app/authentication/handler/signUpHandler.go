package handler

import (
	"authentication/business"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type signUpController struct {
	service *business.SignUpService
}

func NewSignUpController(service *business.SignUpService) *signUpController {
	return &signUpController{
		service: service,
	}
}

func (controller *signUpController) SignUp(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(user); err != nil {
		validationErrors := validations.FormatValidationErrors(ctx.Request.Context(), err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, gin.H{
			genericConstants.GenericJSONErrorMessage: genericConstants.ValidatorError,
			genericConstants.GenericValidationError:  validationErrors,
		})
		return
	}
	if err := controller.service.SignUp(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

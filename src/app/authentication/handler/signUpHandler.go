package handler

import (
	"authentication/business"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"
)

type signUpController struct {
	service *business.SignUpService
}

func NewSignUpController(service *business.SignUpService) *signUpController {
	return &signUpController{
		service: service,
	}
}

// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags SignUP
// @Accept json
// @Produce json
// @Param user body models.UserSignUp true "User details"
// @Success 200 {string} string "User created successfully"
// @Failure 400 {string} string "Bad request"
// @Router /v1/auth/signup [post]
func (controller *signUpController) SignUp(ctx *gin.Context) {
	var user models.UserSignUp
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

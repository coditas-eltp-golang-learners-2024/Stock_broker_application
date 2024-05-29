package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"encoding/json"
	genericConstants "stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
	"stock_broker_application/src/utils"
	"stock_broker_application/src/utils/validations"
	"strings"

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

// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags SignUP
// @Accept json
// @Produce json
// @Param user body models.UserSignUp true "User details"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {string} string "Bad request"
// @Router /v1/auth/signup [post]
func (controller *signUpController) SignUp(ctx *gin.Context) {
	var user models.UserSignUp
	if err := ctx.ShouldBindJSON(&user); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		utils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsgs})
		return
	}

	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(user); err != nil {
		validationErrors := validations.FormatValidationErrors(ctx.Request.Context(), err.(validator.ValidationErrors))
		utils.SendBadRequest(ctx, validationErrors)
		return
	}
	if err := controller.service.SignUp(&user); err != nil {
		if strings.Contains(err.Error(), constants.ErrorUserExists) {
			utils.SendConflictError(ctx, err.Error())
		} else {
			utils.SendInternalServerError(ctx, err.Error())
		}

		return
	}

	utils.SendCreated(ctx, constants.SignUpSuccessMessage)
}

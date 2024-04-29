package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"log"
	"net/http"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
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
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(user); err != nil {
		log.Println("Validation error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidPasswordFormat})
		return
	}
	if err := controller.service.SignUp(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

package router

import (
	"authentication/business"
	"authentication/commons/constants"
	serviceConstant "authentication/commons/constants"
	"authentication/handler"
	"authentication/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	//Dependency Injection for forgot-Password-Feature
	repository := repositories.NewForgotPasswordRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewUsersService(repository)
	newUsersController := handler.NewUsersController(service)

	//Dependency Injection for OTP-Validation-Feature
	connectionWithDb := postgres.GetPostGresClient().GormDb
	userRepository := repositories.NewUserRepository(connectionWithDb)
	otpService := business.NewOTPService(userRepository)
	otpValidationController := handler.NewOTPValidationController(otpService)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})

		v1Routes.POST(constants.ValidateOTP, otpValidationController.HandleOTPValidation)
		v1Routes.POST(serviceConstant.ForgotPassword, newUsersController.HandleForgotPassword)

	}
	return router
}

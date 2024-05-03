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

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		connectionWithDb := postgres.GetPostGresClient().GormDb
		userRepository := repositories.NewUserRepository(connectionWithDb)
		otpService := business.NewOTPService(userRepository)

		// v1Routes.POST(constants.ValidateOTP, handler.NewOTPValidationController(otpService))

		otpValidationController := handler.NewOTPValidationController(otpService)
		v1Routes.POST(constants.ValidateOTP, otpValidationController.HandleOTPValidation)

	}
	return router
}

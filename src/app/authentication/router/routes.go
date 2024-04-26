package router

import (
	"net/http"

	"authentication/business"
	serviceConstant "authentication/commons/constants"
	"authentication/handler"
	"authentication/repositories"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
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
		//Add your routes here
		// Create instances of the repository and service layers
		connectionWithDb := postgres.GetPostGresClient().GormDb
		userRepository := repositories.NewCustomerRepository(connectionWithDb)
		otpService := business.NewOTPService(userRepository)

		v1Routes.POST("/validateOTP", handler.NewValidateOTPHandler(otpService))

	}
	return router
}

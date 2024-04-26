package router

import (
	"net/http"

	"authentication/business"
	serviceConstant "authentication/commons/constants"
	"authentication/handler"

	"authentication/repositories"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	//Dependency Injection for Forget-Password-Feature
	repository := repositories.NewForgetPasswordRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewCredentialsService(repository)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		//Add your routes here
		v1Routes.POST(serviceConstant.ForgetPasswordFeature, handler.UpdateCredentialsHandler(service))

	}
	return router
}

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	// "stock_broker_application/src/utils/postgres"
	serviceConstant "watchlist/commons/constants"
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
		// ConnectionWithDb := postgres.GetPostGresClient().GormDb
		// docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstant.DocsAnyPath, ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}

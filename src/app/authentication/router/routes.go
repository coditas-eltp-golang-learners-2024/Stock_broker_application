package router

import (
	"authentication/business"
	serviceConstant "authentication/commons/constants"
	"authentication/docs"
	"authentication/handler"
	"authentication/repositories"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		// // Swagger documentation setup
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstant.DocsAnyPath, ginSwagger.WrapHandler(swaggerFiles.Handler))

		// routes
		connectionWithDb := postgres.GetPostGresClient().GormDb
		userDatabaseRepo := repositories.NewSignInRepositoryImpl(connectionWithDb)
		userService := business.NewSignInService(userDatabaseRepo)

		userRepository := repositories.NewUserRepository(connectionWithDb)
		otpService := business.NewOTPService(userRepository)

		v1Routes.POST(serviceConstant.SignIn, handler.SignInHandler(userService, otpService))
	}
	return router
}

package router

import (
	"authentication/business"
	"authentication/docs"
	serviceConstant "authentication/commons/constants"
	"authentication/handler"
	"authentication/repositories"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
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
	//Dependency Injection for change-password
	connectionWithDb := postgres.GetPostGresClient().GormDb
	userDatabaseRepo := repositories.NewUserDBRepository(connectionWithDb)
	passwordService := business.NewRestPasswordService(userDatabaseRepo)
	changePasswordHandler := handler.NewChangePasswordController(passwordService)
	//Dependency Injection for forgot-Password-Feature
	repository := repositories.NewForgotPasswordRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewUsersService(repository)
	newUsersController := handler.NewUsersController(service)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		v1Routes.PATCH(serviceConstant.ChangePassword, headerCheck.AuthMiddleware(), handler.HandleChangePassword(changePasswordHandler))
		v1Routes.POST(serviceConstant.ForgotPassword, newUsersController.HandleForgotPassword)
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
	return router
}

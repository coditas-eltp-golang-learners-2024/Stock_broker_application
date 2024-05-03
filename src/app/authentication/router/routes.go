package router

import (
	"net/http"

	"authentication/business"
	"authentication/commons/constants"
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

	userRepoSignUp := repositories.NewUserSignUpInstance()
	userSignUpService := business.NewSignUpService(userRepoSignUp)
	userSignUpController := handler.NewSignUpController(userSignUpService)

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
		//Add your routes here
		v1Routes.POST(constants.SignUp, userSignUpController.SignUp)
		v1Routes.POST(serviceConstant.ForgotPassword, newUsersController.HandleForgotPassword)
	}
	return router
}

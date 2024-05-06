package router

import (
	"authentication/business"
	"authentication/commons/constants"
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

func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	userRepoSignUp := repositories.NewUserSignUpInstance()
	userSignUpService := business.NewSignUpService(userRepoSignUp)
	userSignUpController := handler.NewSignUpController(userSignUpService)

	//dependency injection for signin
	connectionWithDb := postgres.GetPostGresClient().GormDb
	userDatabaseRepository := repositories.NewSignInRepository(connectionWithDb)
	signInService := business.NewSignInService(userDatabaseRepository)
	signInController := handler.NewSignInController(signInService)

	//Dependency Injection for forgot-Password-Feature
	repository := repositories.NewForgotPasswordRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewUsersService(repository)
	newUsersController := handler.NewUsersController(service)

	//Dependency Injection for OTP-Validation-Feature
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
		//Add your routes here

		//  Swagger documentation setup
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

		// routes
		v1Routes.POST(constants.SignUp, userSignUpController.SignUp)
		v1Routes.POST(serviceConstant.SignIn, signInController.HandleSignIn)
		v1Routes.POST(serviceConstant.ForgotPassword, newUsersController.HandleForgotPassword)
		v1Routes.POST(constants.ValidateOTP, otpValidationController.HandleValidateOTP)

	}
	return router
}

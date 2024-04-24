package router

import (
	"authentication/business"
	"authentication/constants"
	"authentication/docs"
	"authentication/handlers"
	"authentication/repositories"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter sets up the routes for the application
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Initialize UserRepository
	userRepository := repositories.NewSignInRepositoryImpl(db)

	// Initialize SignInService with UserRepository
	userAuthService := business.NewSignInService(userRepository)

	// Set up routes
	r.POST(constants.SignInRoute, handlers.SignInHandler(userAuthService))

	// Swagger documentation setup
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

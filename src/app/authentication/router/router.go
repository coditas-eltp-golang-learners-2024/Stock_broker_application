// Package router provides functions to set up and configure the application's HTTP router.
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"authentication/business"
	"authentication/docs"
	"authentication/handlers"
	"authentication/repositories"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create a new Gin router with default middleware
	router := gin.Default()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create instances of the repository and service layers
	userRepository := repositories.NewCustomerRepository(db)
	otpService := business.NewOTPService(userRepository)

	router.POST("/validateOTP", handlers.NewValidateOTPHandler(otpService))

	return router
}

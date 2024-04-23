package router

import (
	"authentication/business"
	"authentication/constants"
	"authentication/handlers"
	"authentication/repositories"
	"authentication/utils/database"
	"log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// router is created with gin
func InitializeRouter() *gin.Engine {
	userRouter := gin.New()
	userRouter.Use(gin.Logger()) // LOG DETAILS ON CONSOLE
	ConnectionWithDb, err := database.ConnectionWithDb()
	// checking err while connectingDB
	if err != nil {
		log.Fatalf("%s :%s", constants.ErrConnectingDB.Error(), err.Error())
	}
	userDatabaseRepo := repositories.NewUserDBRepository(ConnectionWithDb)
	passwordService := business.NewRestPasswordService(userDatabaseRepo)
	userRouter.PATCH(constants.CustomerchangepasswordEndpoint, handlers.AuthMiddleware(), handlers.HandleChangePassword(passwordService))
	userRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swaggerAdded
	return userRouter
}

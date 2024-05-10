package router

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	"watchlist/business"
	serviceConstant "watchlist/commons/constants"
	"watchlist/docs"
	"watchlist/handler"
	"watchlist/repositories"
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

	//dependency injection for getWatchlistScrips
	connectionWithDb := postgres.GetPostGresClient().GormDb
	userDatabaseRepository := repositories.NewWatchlistRepository(connectionWithDb)
	watchlistService := business.NewWatchlistScripsService(userDatabaseRepository)
	watchlistController := handler.NewWatchlistController(watchlistService)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{

		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1Routes.GET(serviceConstant.GetWatchList, headerCheck.AuthMiddleware(), watchlistController.HandleWatchlistScrips)

	}
	return router
}

package router

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	"watchlist/business"
	serviceConstant "watchlist/commons/constants"
	"watchlist/handler"
	"watchlist/repositories"

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

	//dependency injection for signin
	connectionWithDb := postgres.GetPostGresClient().GormDb
	userDatabaseRepository := repositories.NewWatchlistRepository(connectionWithDb)
	watchlistService := business.NewWatchlistService(userDatabaseRepository)
	watchlistController := handler.NewWatchlistController(watchlistService)

	

	v2Routes := router.Group(genericConstants.RouterV2Config)
	{
	
		// routes
		v2Routes.GET(serviceConstant.GetWatchList, headerCheck.AuthMiddleware(), watchlistController.HandleWatchlist)
		
	}
	return router
}

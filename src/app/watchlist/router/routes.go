package router

import (
	"watchlist/business"
	"watchlist/docs"
	"watchlist/handler"
	"watchlist/repositories"

	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	serviceConstant "watchlist/commons/constants"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	//Dependency injection for getWatchlist-feature
	repository := repositories.NewGetWatclistsRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewUsersService(repository)
	newGetwatchlistController := handler.NewGetWatchListController(service)

	v2Routes := router.Group(genericConstants.RouterV2Config)
	{
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v2Routes.GET(serviceConstant.GetWatchList, headerCheck.AuthMiddleware(), newGetwatchlistController.HandleGetWatchlist)
		v2Routes.GET(serviceConstant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}

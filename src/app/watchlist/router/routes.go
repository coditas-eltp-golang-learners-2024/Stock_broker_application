package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	"watchlist/business"
	serviceConstants "watchlist/commons/constants"
	"watchlist/docs"
	"watchlist/handler"
	"watchlist/repositories"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())
	connectionWithDb := postgres.GetPostGresClient().GormDb

	userDatabaseRepository := repositories.NewUserDBRepository(connectionWithDb)
	createWatchlistHandler := business.NewCreateWatchlistService(userDatabaseRepository)
	createWatchlistController := handler.NewWatchlistController(createWatchlistHandler)

	repository := repositories.NewGetWatclistsRepository(connectionWithDb)
	service := business.NewUsersService(repository)
	newGetwatchlistsController := handler.NewGetWatchListsController(service)

	deleteWatchlistRepository := repositories.NewDeleteWatchlistRepository(connectionWithDb)
	deleteWatchlistScripsService := business.NewDeleteWatchlistService(deleteWatchlistRepository)
	deleteWatchlistScripsController := handler.NewDeleteWatchlistScripsController(deleteWatchlistScripsService)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstants.WatchlistHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		v1Routes.GET(serviceConstants.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1Routes.POST(serviceConstants.CreateWatchlist, headerCheck.AuthMiddleware(), createWatchlistController.HandleCreateWatchlist)
		v1Routes.GET(serviceConstants.GetWatchLists, headerCheck.AuthMiddleware(), newGetwatchlistsController.HandleGetWatchlists)
		v1Routes.DELETE(serviceConstants.DeleteWatchlistScrips, headerCheck.AuthMiddleware(), deleteWatchlistScripsController.HandleDeleteWatchlistScrips)
	}
	return router
}

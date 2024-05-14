package router

import (
	"authentication/docs"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"watchlist/business"
	"watchlist/commons/constants"
	"watchlist/handler"
	"watchlist/repositories"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"stock_broker_application/src/utils/postgres"
	serviceConstants "watchlist/commons/constants"

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

	connectionWithDb := postgres.GetPostGresClient().GormDb
	userDatabaseRepository := repositories.NewUserDBRepository(connectionWithDb)
	createWatchlistHandler := business.NewCreateWatchlistService(userDatabaseRepository)
	createWatchlistController := handler.NewWatchlistController(createWatchlistHandler)

	editWatchList := repositories.NewRenameWatchListRepository()
	editWatchListService := business.NewRenameWatchListService(editWatchList)
	editWatchListController := handler.NewRenameWatchListController(editWatchListService)

	deleteWatchList := repositories.NewDeleteWatchListRepository()
	deleteWatchListService := business.NewDeleteWatchListService(deleteWatchList)
	deleteWatchListController := handler.NewDeleteWatchListDeleteController(deleteWatchListService)

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
		v1Routes.PUT(constants.RenameWatchList, headerCheck.AuthMiddleware(), editWatchListController.EditWatchList)
		v1Routes.DELETE(constants.DeleteWatchList, headerCheck.AuthMiddleware(), deleteWatchListController.DeleteWatchList)
	}
	return router
}

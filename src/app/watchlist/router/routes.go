package router

import (
	"net/http"
	"watchlist/business"
	"watchlist/handler"
	"watchlist/repositories"

	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	serviceConstant "watchlist/commons/constants"

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

	//Dependency injection for getWatchlist-feature
	repository := repositories.NewGetWatclistsRepository(postgres.GetPostGresClient().GormDb)
	service := business.NewUsersService(repository)
	newGetwatchlistController := handler.NewGetWatchListController(service)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		v1Routes.GET(serviceConstant.GetWatchList, headerCheck.AuthMiddleware(), newGetwatchlistController.HandleGetWatchlist)
	}
	return router
}

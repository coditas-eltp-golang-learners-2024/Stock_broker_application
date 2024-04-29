package router

import (
	"net/http"

	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/commons/constants"
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

	editWatchList := repositories.NewEditWatchListRepository()
	editWatchListService := business.NewEditWatchListService(editWatchList)
	editWatchListController := handler.NewEditWatchListController(editWatchListService)

	deleteWatchList := repositories.NewDeleteWatchListRepository()
	deleteWatchListService := business.NewDeleteWatchListService(deleteWatchList)
	deleteWatchListController := handler.NewEditWatchListDeleteController(deleteWatchListService)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		//Add your routes here
		v1Routes.POST(constants.EditWatchList, editWatchListController.EditWatchList)
		v1Routes.DELETE(constants.DeleteWatchList, deleteWatchListController.DeleteWatchList)
	}
	return router
}

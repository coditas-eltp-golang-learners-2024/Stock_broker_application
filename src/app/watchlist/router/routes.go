package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/middleware/headerCheck"
	"stock_broker_application/src/utils/postgres"
	"watchlist/business"
	serviceConstant "watchlist/commons/constants"
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
	deleteWatchlistRepository := repositories.NewDeleteWatchlistRepository(connectionWithDb)
	deleteWatchlistScripsService := business.NewDeleteWatchlistService(deleteWatchlistRepository)
	deleteWatchlistScripsHandler := handler.NewDeleteWatchlistScripsHandler(*deleteWatchlistScripsService)

	// Apply authentication middleware
	router.Use(headerCheck.AuthMiddleware())

	v1Routes := router.Group(constants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.AuthenticationHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				constants.ResponseMessageKey: constants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})

		v1Routes.DELETE(serviceConstant.DeleteWatchlistScrips, headerCheck.AuthMiddleware(), deleteWatchlistScripsHandler.HandleDeleteWatchlistScrips)
	}
	return router
}

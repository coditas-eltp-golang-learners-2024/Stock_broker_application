package handler

import (
	"net/http"

	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/commons/constants"

	"github.com/gin-gonic/gin"
)

type NewGetWatchlistController interface {
	HandleGetWatchlist(ctx *gin.Context)
}

type getWatchlistController struct {
	service business.NewGetWatchlistService
}

func NewGetWatchListController(service business.NewGetWatchlistService) NewGetWatchlistController {
	return &getWatchlistController{
		service: service,
	}
}

func (controller *getWatchlistController) HandleGetWatchlist(context *gin.Context) {

	watchlistData, err := controller.service.NewGetWatchlistService(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.WatchlistNotFound})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: watchlistData})

}

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

// @Summary Retrieve user's watchlist
// @Description Handler function to fetch the user's watchlist data.
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns the user's watchlist data"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /get-watchlist [get]
func (controller *getWatchlistController) HandleGetWatchlist(context *gin.Context) {

	watchlistData, err := controller.service.NewGetWatchlistService(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.WatchlistNotFound})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.Watchlist: watchlistData})

}

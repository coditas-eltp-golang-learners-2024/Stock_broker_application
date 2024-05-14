package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
)

type watchlistScripsController struct {
	service *business.WatchlistScripsService
}

func NewWatchlistScripController(service *business.WatchlistScripsService) *watchlistScripsController {
	return &watchlistScripsController{
		service: service,
	}
}

// HandleWatchlistScrips handles the request to get watchlist scrips.
// @Summary Get watchlist scrips
// @Description Retrieves the list of scrips for a given watchlist
// @Tags Get Watchlist Scrips
// @Accept json
// @Produce json
// @Security JWT
// @Param watchlist_name query string true "Name of the watchlist"
// @Success 200 {object} models.Scrip "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /v1/watchlist/scrips/list [get]
func (controller *watchlistScripsController) HandleWatchlistScrips(context *gin.Context) {
	watchlistName := context.Query(genericConstants.WatchlistName)
	if watchlistName == "" {
		context.JSON(http.StatusNoContent, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.WatchlistNameRequiredError})
		return
	}
	scrips, err := controller.service.GetScrips(context, watchlistName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.ScripsKey: scrips})
}

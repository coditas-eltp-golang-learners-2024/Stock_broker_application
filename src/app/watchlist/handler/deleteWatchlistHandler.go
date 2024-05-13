package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/commons/constants"
)

type DeleteWatchlistScripsHandler struct {
	deleteWatchlistService business.DeleteWatchlistService
}

func NewDeleteWatchlistScripsHandler(service business.DeleteWatchlistService) *DeleteWatchlistScripsHandler {
	return &DeleteWatchlistScripsHandler{
		deleteWatchlistService: service,
	}
}

// HandleDeleteWatchlistScrips handles HTTP DELETE requests for deleting scrips from a watchlist.
// @Summary Delete scrips from a watchlist
// @Description Delete scrips from a watchlist based on the provided watchlist name and scrips
// @ID delete-watchlist-scrips
// @Accept  json
// @Produce  json
// @Param request body DeleteWatchlistScripsRequest true "Delete Watchlist Scrips Request"
// @Success 200 {object} gin.H{"message": "Watchlist scrips deleted successfully"}
// @Failure 400 {object} gin.H{"message": "Invalid request payload"}
// @Failure 500 {object} gin.H{"message": "Failed to delete watchlist scrips"}
// @Router /watchlist/scrips [delete]
func (controller *DeleteWatchlistScripsHandler) HandleDeleteWatchlistScrips(context *gin.Context) {
	request := struct {
		WatchlistName string   `json:"watchlist_name"`
		Scrips        []string `json:"scrips"`
	}{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONMessage: constants.ErrorInvalidRequestPayload})
		return
	}

	// Call the service to delete scrips from the watchlist
	err := controller.deleteWatchlistService.DeleteScripsFromWatchlist(context, request.WatchlistName, request.Scrips)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONMessage: constants.ErrorFailedToDeleteScrips})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.DeletedWatchlistScrips})
}

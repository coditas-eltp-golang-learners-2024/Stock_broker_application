package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/commons/constants"
	"watchlist/models"
)

type DeleteWatchlistScripsHandler struct {
	deleteWatchlistService *business.DeleteWatchlistService
}

func NewDeleteWatchlistScripsHandler(service *business.DeleteWatchlistService) *DeleteWatchlistScripsHandler {
	return &DeleteWatchlistScripsHandler{
		deleteWatchlistService: service,
	}
}

// HandleDeleteWatchlistScrips handles HTTP DELETE requests for deleting scrips from a watchlist.
// @Summary Delete scrips from a watchlist
// @Description Delete scrips from a watchlist based on the provided watchlist name and scrips
// @Accept  json
// @Produce  json
// @Param request body models.DeleteWatchlistScripsRequest true "Delete Watchlist Scrips Request"
// @Success 200 {string} string "Watchlist scrips deleted successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Router /v1/watchlist/scrips [delete]
func (controller *DeleteWatchlistScripsHandler) HandleDeleteWatchlistScrips(context *gin.Context) {
	var request models.DeleteWatchlistScripsRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONMessage: constants.InvalidRequestPayloadError})
		return
	}

	// Call the service to delete scrips from the watchlist
	err := controller.deleteWatchlistService.DeleteScripsFromWatchlist(context, request.WatchlistName, request.Scrips)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONMessage: constants.FailedToDeleteScripsError})
		return
	}
	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.DeletedWatchlistScrips})
}

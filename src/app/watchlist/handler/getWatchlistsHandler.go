package handler

import (
	"errors"
	"net/http"
	"stock_broker_application/src/utils"

	"watchlist/business"

	"github.com/gin-gonic/gin"
)

// NewGetWatchlistController defines the interface for handling GetWatchlist requests.
type NewGetWatchlistController interface {
	HandleGetWatchlists(ctx *gin.Context)
}

type getWatchlistController struct {
	service business.NewGetWatchlistsService
}

// NewGetWatchListsController creates a new instance of GetWatchlistController.
func NewGetWatchListsController(service business.NewGetWatchlistsService) NewGetWatchlistController {
	return &getWatchlistController{
		service: service,
	}
}

// HandleGetWatchlists handles the GetWatchlists request.
// @Summary Get the list of WatchLists
// @Description Handler function to fetch the user's watchlist data.
// @Produce json
// @Tags GetWatchlists
// @Success 200 {object} models.GetWatchlists "Returns the user's watchlist data"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security JWT
// @Router /v1/watchlist/list [get]
func (controller *getWatchlistController) HandleGetWatchlists(context *gin.Context) {

	watchlistData, err := controller.service.NewGetWatchlistsService(context)

	if err != nil {
		utils.SendInternalServerError(context, err.Error())
		return
	}

	if len(watchlistData.Watchlist) == 0 {
		err := errors.New(http.StatusText(http.StatusNoContent))
		utils.SendNoContentError(context, err)
		return
	}

	utils.SendStatusOk(context, watchlistData)
}

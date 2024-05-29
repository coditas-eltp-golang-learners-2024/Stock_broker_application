package handler

import (
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	serviceConstants "watchlist/commons/constants"
	serviceModel "watchlist/models"

	"github.com/gin-gonic/gin"
)

type CreateWatchlistController struct {
	service *business.CreateWatchlistService
}

func NewWatchlistController(service *business.CreateWatchlistService) *CreateWatchlistController {
	return &CreateWatchlistController{
		service: service,
	}
}

// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags CreateWatchlist
// @Accept json
// @Produce json
// @Security JWT
// @Param user body models.CreateWatchlist true "User details"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {string} string "Bad request"
// @Router /v1/watchlist/create [post]
func (controller *CreateWatchlistController) HandleCreateWatchlist(ctx *gin.Context) {
	var createWatchlistRequest serviceModel.CreateWatchlist
	if err := ctx.BindJSON(&createWatchlistRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	if err := controller.service.CreateWatchlistService(createWatchlistRequest, ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{genericConstants.AccessToken: err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{genericConstants.BFFResponseSuccessMessage: serviceConstants.WatchlistCreatedSuccessMessage})
}

package handler

import (
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/validations"
	"watchlist/business"
	constants "watchlist/commons/constants"
	"watchlist/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type renameWatchListController struct {
	service *business.RenameWatchListService
}

func NewRenameWatchListController(service *business.RenameWatchListService) *renameWatchListController {
	return &renameWatchListController{
		service: service,
	}
}

// @Summary Edit a watchlist
// @Description Edit a watchlist with the provided details
// @Tags EditWatchlist
// @Accept json
// @Produce json
// @Security JWT
// @Param watchlist body models.WatchlistRenameModel true "Watchlist details"
// @Success 200 {string} string "Watchlist edited successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/edit-watchlist [put]
func (controller *renameWatchListController) EditWatchList(ctx *gin.Context) {
	var watchlist models.WatchlistRenameModel
	if err := ctx.ShouldBindJSON(&watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(watchlist); err != nil {
		validationErrors := validations.FormatValidationErrors(ctx.Request.Context(), err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, gin.H{
			genericConstants.GenericJSONErrorMessage: err.Error(),
			genericConstants.GenericValidationError:  validationErrors,
		})
		return
	}
	if err := controller.service.RenameWatchList(&watchlist, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: constants.InternalServerError})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{genericConstants.BFFResponseSuccessMessage: constants.WatchlistRenameSuccessMessage})
}

package business

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type RenameWatchListService struct {
	RenameWatchListRepository repositories.RenameWatchListRepository
}

func NewRenameWatchListService(renameWatchListRepository repositories.RenameWatchListRepository) *RenameWatchListService {
	return &RenameWatchListService{
		RenameWatchListRepository: renameWatchListRepository,
	}
}

func (service *RenameWatchListService) RenameWatchList(watchlist *models.RenameWatchlist, ctx *gin.Context) error {
	client := postgres.GetPostGresClient()
	userId := ctx.Value(genericConstants.Id).(uint16)
	err := service.RenameWatchListRepository.RenameWatchlist(client.GormDb, watchlist, userId)
	if err != nil {
		return err
	}
	return nil

}

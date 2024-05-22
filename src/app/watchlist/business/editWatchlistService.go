package business

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type EditWatchlistService struct {
	EditWatchlistRepository repositories.EditWatchlistRepository
}

func NewEditWatchlistService(editWatchlistRepository repositories.EditWatchlistRepository) *EditWatchlistService {
	return &EditWatchlistService{
		EditWatchlistRepository: editWatchlistRepository,
	}
}

func (service *EditWatchlistService) EditWatchlist(editWatchlist *models.EditWatchlistRequest, ctx *gin.Context) error {

	client := postgres.GetPostGresClient()
	userId := ctx.Value(genericConstants.Id).(uint16)
	err := service.EditWatchlistRepository.InsertScripsToWatchlist(*editWatchlist, userId, client.GormDb)
	if err != nil {
		return err
	}
	return nil
}

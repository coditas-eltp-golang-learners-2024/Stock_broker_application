package business

import (
	genericConstants "stock_broker_application/src/constants"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type NewGetWatchlistsService interface {
	NewGetWatchlistsService(ctx *gin.Context) ([]string, error)
}

type getWatchListSercvice struct {
	getWatchlistsRepository repositories.GetWatchlistsRepository
}

func NewUsersService(getWatchlistRepository repositories.GetWatchlistsRepository) NewGetWatchlistsService {
	return &getWatchListSercvice{
		getWatchlistsRepository: getWatchlistRepository,
	}
}

func (service *getWatchListSercvice) NewGetWatchlistsService(ctx *gin.Context) ([]string, error) {
	id := ctx.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId: id,
	}
	return service.getWatchlistsRepository.GetWatchlists(ctx, condition)
}

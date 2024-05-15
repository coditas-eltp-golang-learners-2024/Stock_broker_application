package business

import (
	genericConstants "stock_broker_application/src/constants"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type NewGetWatchlistsService interface {
	NewGetWatchlistsService(ctx *gin.Context) ([]string, error)
}

type getWatchListsSercvice struct {
	getWatchlistsRepository repositories.GetWatchlistsRepository
}

func NewUsersService(getWatchlistInstance repositories.GetWatchlistsRepository) NewGetWatchlistsService {
	return &getWatchListsSercvice{
		getWatchlistsRepository: getWatchlistInstance,
	}
}

func (service *getWatchListsSercvice) NewGetWatchlistsService(ctx *gin.Context) ([]string, error) {
	id := ctx.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId: id,
	}
	return service.getWatchlistsRepository.GetWatchlists(ctx, condition)
}

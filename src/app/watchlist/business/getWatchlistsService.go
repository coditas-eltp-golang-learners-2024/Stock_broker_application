package business

import (
	genericConstants "stock_broker_application/src/constants"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type NewGetWatchlistsService interface {
	NewGetWatchlistsService(ctx *gin.Context) (models.GetWatchlists, error)
}

type getWatchListsSercvice struct {
	getWatchlistsRepository repositories.GetWatchlistsRepository
}

func NewUsersService(getWatchlistsInstance repositories.GetWatchlistsRepository) NewGetWatchlistsService {
	return &getWatchListsSercvice{
		getWatchlistsRepository: getWatchlistsInstance,
	}
}

func (service *getWatchListsSercvice) NewGetWatchlistsService(ctx *gin.Context) (models.GetWatchlists, error) {
	id := ctx.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId: id,
	}
	return service.getWatchlistsRepository.GetWatchlists(ctx, condition)
}
